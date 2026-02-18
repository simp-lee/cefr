package cefr

import (
	"math"
	"strings"
	"testing"
)

func TestShouldSample(t *testing.T) {
	tests := []struct {
		name      string
		wordCount int
		threshold int
		want      bool
	}{
		{"below threshold", 5000, 10000, false},
		{"at threshold", 10000, 10000, false},
		{"above threshold", 10001, 10000, true},
		{"zero words", 0, 10000, false},
		{"custom threshold", 500, 300, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shouldSample(tt.wordCount, tt.threshold)
			if got != tt.want {
				t.Errorf("shouldSample(%d, %d) = %v; want %v",
					tt.wordCount, tt.threshold, got, tt.want)
			}
		})
	}
}

// buildLongText creates a text with approximately n words using complete sentences.
func buildLongText(n int) string {
	// Each sentence has 10 words.
	sentence := "The quick brown fox jumps over the lazy dog daily."
	wordsPerSentence := 10
	sentencesNeeded := (n + wordsPerSentence - 1) / wordsPerSentence

	var b strings.Builder
	for i := 0; i < sentencesNeeded; i++ {
		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(sentence)
	}
	return b.String()
}

func TestSampleText(t *testing.T) {
	t.Run("short text returns single element", func(t *testing.T) {
		text := "This is a short text."
		got := sampleText(text, 5)
		if len(got) != 1 {
			t.Fatalf("expected 1 segment, got %d", len(got))
		}
		if got[0] != text {
			t.Errorf("expected original text, got %q", got[0])
		}
	})

	t.Run("text under 3000 words returns single element", func(t *testing.T) {
		text := buildLongText(2500)
		got := sampleText(text, 2500)
		if len(got) != 1 {
			t.Fatalf("expected 1 segment, got %d", len(got))
		}
	})

	t.Run("long text returns three segments", func(t *testing.T) {
		text := buildLongText(5000)
		got := sampleText(text, 5000)
		if len(got) != 3 {
			t.Fatalf("expected 3 segments, got %d", len(got))
		}
		// Each segment should be non-empty
		for i, seg := range got {
			if strings.TrimSpace(seg) == "" {
				t.Errorf("segment %d is empty", i)
			}
		}
		// Each segment should end with sentence-ending punctuation
		for i, seg := range got {
			trimmed := strings.TrimSpace(seg)
			last := trimmed[len(trimmed)-1]
			if last != '.' && last != '!' && last != '?' {
				t.Errorf("segment %d does not end at sentence boundary: %q",
					i, trimmed[max(0, len(trimmed)-50):])
			}
		}
	})

	t.Run("segments have roughly 1000 words each", func(t *testing.T) {
		text := buildLongText(6000)
		got := sampleText(text, 6000)
		if len(got) != 3 {
			t.Fatalf("expected 3 segments, got %d", len(got))
		}
		for i, seg := range got {
			wc := len(strings.Fields(seg))
			// Allow ~500-1500 range for flexibility at sentence boundaries
			if wc < 500 || wc > 1500 {
				t.Errorf("segment %d has %d words, expected ~1000", i, wc)
			}
		}
	})

	t.Run("when no room for middle segment does not duplicate tail", func(t *testing.T) {
		makeSentence := func(word string, count int) string {
			var b strings.Builder
			for i := 0; i < count; i++ {
				if i > 0 {
					b.WriteByte(' ')
				}
				b.WriteString(word)
			}
			b.WriteByte('.')
			return b.String()
		}

		// 2 huge sentences => long enough to sample, but no separate middle region.
		text := makeSentence("alpha", 1600) + " " + makeSentence("beta", 1600)
		got := sampleText(text, 3200)

		if len(got) < 1 || len(got) > 2 {
			t.Fatalf("expected 1 or 2 unique segments, got %d", len(got))
		}

		seen := make(map[string]bool)
		for i, seg := range got {
			seg = strings.TrimSpace(seg)
			if seg == "" {
				t.Fatalf("segment %d is empty", i)
			}
			if seen[seg] {
				t.Fatalf("segment %d duplicates another segment", i)
			}
			seen[seg] = true
		}
	})
}

func TestMergeSampledResults_EmptyInput(t *testing.T) {
	cfg := defaultConfig()

	t.Run("nil slice", func(t *testing.T) {
		got := mergeSampledResults(nil, cfg)
		if got.Score != 0 {
			t.Errorf("Score = %v; want 0", got.Score)
		}
		if got.Level != "" {
			t.Errorf("Level = %q; want empty", got.Level)
		}
		if got.WordCount != 0 {
			t.Errorf("WordCount = %d; want 0", got.WordCount)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		got := mergeSampledResults([]Result{}, cfg)
		if got.Score != 0 {
			t.Errorf("Score = %v; want 0", got.Score)
		}
		if got.Level != "" {
			t.Errorf("Level = %q; want empty", got.Level)
		}
		if got.WordCount != 0 {
			t.Errorf("WordCount = %d; want 0", got.WordCount)
		}
	})
}

func TestMergeSampledResults_SingleResult(t *testing.T) {
	r := Result{
		Level:      "B1",
		Score:      3.0,
		Confidence: 0.8,
		Vocab: VocabResult{
			Score:        3.0,
			Distribution: map[string]float64{"B1": 0.5, "A2": 0.3, "A1": 0.2},
			UnknownRatio: 0.05,
			ContentWords: 100,
		},
		Syntax: SyntaxResult{
			Score:              3.0,
			AvgSentenceLength:  15.0,
			SubordinationIndex: 0.5,
			PassiveRate:        0.1,
			ConnectorDiversity: 8,
		},
		Readability: ReadabilityResult{
			Score: 3.0,
			FKGL:  7.0,
			FRE:   65.0,
			CLI:   8.0,
		},
		WordCount:     500,
		SentenceCount: 30,
	}
	got := mergeSampledResults([]Result{r}, defaultConfig())
	if got.Score != r.Score {
		t.Errorf("single result: Score = %v; want %v", got.Score, r.Score)
	}
	if got.Confidence != r.Confidence {
		t.Errorf("single result: Confidence = %v; want %v", got.Confidence, r.Confidence)
	}
}

func TestMergeSampledResults_MultipleResults(t *testing.T) {
	results := []Result{
		{
			Level: "A2",
			Score: 2.0,
			Vocab: VocabResult{
				Score:        2.0,
				Distribution: map[string]float64{"A1": 0.4, "A2": 0.6},
				UnknownRatio: 0.02,
				ContentWords: 80,
			},
			Syntax: SyntaxResult{
				Score:              2.0,
				AvgSentenceLength:  10.0,
				SubordinationIndex: 0.3,
				PassiveRate:        0.05,
				ConnectorDiversity: 5,
			},
			Readability: ReadabilityResult{
				Score: 2.0,
				FKGL:  4.0,
				FRE:   80.0,
				CLI:   5.0,
			},
			WordCount:     400,
			SentenceCount: 40,
		},
		{
			Level: "B2",
			Score: 4.0,
			Vocab: VocabResult{
				Score:        4.0,
				Distribution: map[string]float64{"B1": 0.3, "B2": 0.7},
				UnknownRatio: 0.08,
				ContentWords: 120,
			},
			Syntax: SyntaxResult{
				Score:              4.0,
				AvgSentenceLength:  20.0,
				SubordinationIndex: 1.0,
				PassiveRate:        0.2,
				ConnectorDiversity: 15,
			},
			Readability: ReadabilityResult{
				Score: 4.0,
				FKGL:  10.0,
				FRE:   60.0,
				CLI:   11.0,
			},
			WordCount:     600,
			SentenceCount: 30,
		},
		{
			Level: "B1",
			Score: 3.0,
			Vocab: VocabResult{
				Score:        3.0,
				Distribution: map[string]float64{"A2": 0.2, "B1": 0.8},
				UnknownRatio: 0.05,
				ContentWords: 100,
			},
			Syntax: SyntaxResult{
				Score:              3.0,
				AvgSentenceLength:  15.0,
				SubordinationIndex: 0.6,
				PassiveRate:        0.1,
				ConnectorDiversity: 10,
			},
			Readability: ReadabilityResult{
				Score: 3.0,
				FKGL:  7.0,
				FRE:   70.0,
				CLI:   8.0,
			},
			WordCount:     500,
			SentenceCount: 35,
		},
	}

	got := mergeSampledResults(results, defaultConfig())

	// Vocab score: avg(2, 4, 3) = 3.0
	if math.Abs(got.Vocab.Score-3.0) > 0.01 {
		t.Errorf("Vocab.Score = %v; want 3.0", got.Vocab.Score)
	}
	// Syntax score: avg(2, 4, 3) = 3.0
	if math.Abs(got.Syntax.Score-3.0) > 0.01 {
		t.Errorf("Syntax.Score = %v; want 3.0", got.Syntax.Score)
	}
	// Readability score: avg(2, 4, 3) = 3.0
	if math.Abs(got.Readability.Score-3.0) > 0.01 {
		t.Errorf("Readability.Score = %v; want 3.0", got.Readability.Score)
	}

	// WordCount: sum(400, 600, 500) = 1500
	if got.WordCount != 1500 {
		t.Errorf("WordCount = %d; want 1500", got.WordCount)
	}
	// SentenceCount: sum(40, 30, 35) = 105
	if got.SentenceCount != 105 {
		t.Errorf("SentenceCount = %d; want 105", got.SentenceCount)
	}

	// ContentWords: sum(80, 120, 100) = 300
	if got.Vocab.ContentWords != 300 {
		t.Errorf("Vocab.ContentWords = %d; want 300", got.Vocab.ContentWords)
	}

	// UnknownRatio: avg(0.02, 0.08, 0.05) = 0.05
	if math.Abs(got.Vocab.UnknownRatio-0.05) > 0.001 {
		t.Errorf("Vocab.UnknownRatio = %v; want 0.05", got.Vocab.UnknownRatio)
	}

	// Distribution should be averaged
	// A1: avg(0.4, 0, 0) = 0.4/3 ≈ 0.1333
	if math.Abs(got.Vocab.Distribution["A1"]-0.4/3) > 0.01 {
		t.Errorf("Distribution[A1] = %v; want %v", got.Vocab.Distribution["A1"], 0.4/3)
	}

	// Syntax fields: averaged
	if math.Abs(got.Syntax.AvgSentenceLength-15.0) > 0.01 {
		t.Errorf("Syntax.AvgSentenceLength = %v; want 15.0", got.Syntax.AvgSentenceLength)
	}
	if math.Abs(got.Syntax.SubordinationIndex-0.6333) > 0.01 {
		t.Errorf("Syntax.SubordinationIndex = %v; want ~0.633", got.Syntax.SubordinationIndex)
	}
	if math.Abs(got.Syntax.PassiveRate-0.1167) > 0.01 {
		t.Errorf("Syntax.PassiveRate = %v; want ~0.117", got.Syntax.PassiveRate)
	}
	if got.Syntax.ConnectorDiversity != 10 {
		t.Errorf("Syntax.ConnectorDiversity = %d; want 10", got.Syntax.ConnectorDiversity)
	}

	// Readability fields: averaged
	if math.Abs(got.Readability.FKGL-7.0) > 0.01 {
		t.Errorf("Readability.FKGL = %v; want 7.0", got.Readability.FKGL)
	}
	if math.Abs(got.Readability.FRE-70.0) > 0.01 {
		t.Errorf("Readability.FRE = %v; want 70.0", got.Readability.FRE)
	}
	if math.Abs(got.Readability.CLI-8.0) > 0.01 {
		t.Errorf("Readability.CLI = %v; want 8.0", got.Readability.CLI)
	}

	// Score should be recomputed: default weights 0.5*3.0 + 0.3*3.0 + 0.2*3.0 = 3.0
	if math.Abs(got.Score-3.0) > 0.01 {
		t.Errorf("Score = %v; want 3.0", got.Score)
	}
	// Level from score 3.0 → B1
	if got.Level != "B1" {
		t.Errorf("Level = %q; want B1", got.Level)
	}
	// Confidence should be 0 (placeholder for Step 13)
	if got.Confidence != 0 {
		t.Errorf("Confidence = %v; want 0", got.Confidence)
	}
}

func TestSampledVariance(t *testing.T) {
	tests := []struct {
		name    string
		results []Result
		want    float64
	}{
		{
			name:    "nil input",
			results: nil,
			want:    0.0,
		},
		{
			name:    "empty input",
			results: []Result{},
			want:    0.0,
		},
		{
			name:    "single result",
			results: []Result{{Score: 3.0}},
			want:    0.0,
		},
		{
			name: "identical scores",
			results: []Result{
				{Score: 3.0},
				{Score: 3.0},
				{Score: 3.0},
			},
			want: 0.0,
		},
		{
			name: "varied scores",
			results: []Result{
				{Score: 2.0},
				{Score: 4.0},
				{Score: 3.0},
			},
			// mean=3.0, variance = ((2-3)^2 + (4-3)^2 + (3-3)^2) / 3 = 2/3
			want: 2.0 / 3.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sampledVariance(tt.results)
			if math.Abs(got-tt.want) > 0.001 {
				t.Errorf("sampledVariance() = %v; want %v", got, tt.want)
			}
		})
	}
}
