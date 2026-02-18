package cefr

import (
	"math"
	"testing"
)

func syntaxAlmostEqual(a, b, tol float64) bool {
	return math.Abs(a-b) < tol
}

func TestIsRegularPastParticiple(t *testing.T) {
	tests := []struct {
		word string
		want bool
	}{
		{"walked", true},
		{"played", true},
		{"used", true},
		{"running", false},
		{"go", false},
		{"seen", false},
		{"", false},
		{"ed", true},         // edge: just "ed"
		{"interested", true}, // common adjective/past participle
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			if got := isRegularPastParticiple(tt.word); got != tt.want {
				t.Errorf("isRegularPastParticiple(%q) = %v; want %v", tt.word, got, tt.want)
			}
		})
	}
}

func TestCalcAvgSentenceLength(t *testing.T) {
	tests := []struct {
		name          string
		wordCount     int
		sentenceCount int
		want          float64
	}{
		{"normal", 20, 2, 10.0},
		{"single_sentence", 15, 1, 15.0},
		{"zero_sentences", 10, 0, 0.0},
		{"zero_words", 0, 5, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcAvgSentenceLength(tt.wordCount, tt.sentenceCount)
			if !syntaxAlmostEqual(got, tt.want, 0.001) {
				t.Errorf("calcAvgSentenceLength(%d, %d) = %.4f; want %.4f",
					tt.wordCount, tt.sentenceCount, got, tt.want)
			}
		})
	}
}

func TestMapASLToScore(t *testing.T) {
	tests := []struct {
		name string
		asl  float64
		want float64
	}{
		{"below_min", 3.0, 1.0},
		{"at_min", 6.0, 1.0},
		{"at_9", 9.0, 2.0},
		{"at_13", 13.0, 3.0},
		{"at_18", 18.0, 4.0},
		{"at_23", 23.0, 5.0},
		{"at_28", 28.0, 6.0},
		{"above_max", 35.0, 6.0},
		// Midpoint between 6 and 9: asl=7.5 → 1.5
		{"mid_6_9", 7.5, 1.5},
		// Midpoint between 13 and 18: asl=15.5 → 3.5
		{"mid_13_18", 15.5, 3.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapASLToScore(tt.asl)
			if !syntaxAlmostEqual(got, tt.want, 0.01) {
				t.Errorf("mapASLToScore(%.1f) = %.4f; want %.4f", tt.asl, got, tt.want)
			}
		})
	}
}

func TestCalcSubordinationIndex(t *testing.T) {
	tests := []struct {
		name          string
		tokens        []Token
		sentenceCount int
		want          float64
	}{
		{
			"no_subordinators",
			[]Token{
				{Lower: "the"},
				{Lower: "cat"},
				{Lower: "sat"},
			},
			1, 0.0,
		},
		{
			"one_subordinator_one_sentence",
			[]Token{
				{Lower: "she"},
				{Lower: "left"},
				{Lower: "because"},
				{Lower: "it"},
				{Lower: "rained"},
			},
			1, 1.0,
		},
		{
			"two_subordinators_two_sentences",
			[]Token{
				{Lower: "because"},
				{Lower: "he"},
				{Lower: "ran"},
				{Lower: "although"},
				{Lower: "tired"},
			},
			2, 1.0,
		},
		{
			"zero_sentences",
			[]Token{{Lower: "because"}},
			0, 0.0,
		},
		{
			"filtered_subordinator_ignored",
			[]Token{
				{Lower: "because", IsFiltered: true},
				{Lower: "hello"},
			},
			1, 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcSubordinationIndex(tt.tokens, tt.sentenceCount)
			if !syntaxAlmostEqual(got, tt.want, 0.001) {
				t.Errorf("calcSubordinationIndex() = %.4f; want %.4f", got, tt.want)
			}
		})
	}
}

func TestMapSubordinationToScore(t *testing.T) {
	tests := []struct {
		name  string
		index float64
		want  float64
	}{
		{"zero", 0.0, 1.0},
		{"at_0.3", 0.3, 2.0},
		{"at_0.6", 0.6, 3.0},
		{"at_1.0", 1.0, 4.0},
		{"at_1.5", 1.5, 5.0},
		{"at_2.0", 2.0, 6.0},
		{"above_max", 3.0, 6.0},
		{"mid_0_0.3", 0.15, 1.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapSubordinationToScore(tt.index)
			if !syntaxAlmostEqual(got, tt.want, 0.01) {
				t.Errorf("mapSubordinationToScore(%.2f) = %.4f; want %.4f", tt.index, got, tt.want)
			}
		})
	}
}

func TestCalcPassiveRate(t *testing.T) {
	tests := []struct {
		name      string
		sentences []string
		want      float64
	}{
		{
			"no_passive",
			[]string{"The cat sat on the mat."},
			0.0,
		},
		{
			"one_passive_one_sentence",
			[]string{"The book was written by the author."},
			1.0,
		},
		{
			"one_passive_two_sentences",
			[]string{
				"The book was written by the author.",
				"She reads every day.",
			},
			0.5,
		},
		{
			"irregular_past_participle",
			[]string{"The window was broken."},
			1.0,
		},
		{
			"no_sentences",
			[]string{},
			0.0,
		},
		{
			"progressive_not_passive",
			[]string{"She was running quickly."},
			0.0, // "running" ends in -ing, not past participle
		},
		{
			"multiple_passives_one_sentence",
			[]string{"The cake was baked and the house was cleaned."},
			2.0, // two passives / 1 sentence
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcPassiveRate(tt.sentences)
			if !syntaxAlmostEqual(got, tt.want, 0.01) {
				t.Errorf("calcPassiveRate() = %.4f; want %.4f", got, tt.want)
			}
		})
	}
}

func TestMapPassiveRateToScore(t *testing.T) {
	tests := []struct {
		name string
		rate float64
		want float64
	}{
		{"zero", 0.0, 1.0},
		{"at_0.05", 0.05, 2.0},
		{"at_0.10", 0.10, 3.0},
		{"at_0.20", 0.20, 4.0},
		{"at_0.30", 0.30, 5.0},
		{"at_0.40", 0.40, 6.0},
		{"above_max", 0.60, 6.0},
		{"mid_0_0.05", 0.025, 1.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapPassiveRateToScore(tt.rate)
			if !syntaxAlmostEqual(got, tt.want, 0.01) {
				t.Errorf("mapPassiveRateToScore(%.2f) = %.4f; want %.4f", tt.rate, got, tt.want)
			}
		})
	}
}

func TestCalcConnectorDiversity(t *testing.T) {
	tests := []struct {
		name   string
		tokens []Token
		want   int
	}{
		{
			"no_connectors",
			[]Token{{Lower: "cat"}, {Lower: "dog"}},
			0,
		},
		{
			"single_connector",
			[]Token{{Lower: "and"}, {Lower: "cat"}},
			1,
		},
		{
			"multiple_same_connector",
			[]Token{{Lower: "and"}, {Lower: "and"}, {Lower: "cat"}},
			1, // deduplicated
		},
		{
			"multiple_distinct_connectors",
			[]Token{
				{Lower: "and"},
				{Lower: "however"},
				{Lower: "because"},
				{Lower: "first"},
				{Lower: "overall"},
			},
			5, // one from each category
		},
		{
			"multiple_same_category_connectors",
			[]Token{
				{Lower: "and"},
				{Lower: "moreover"},
				{Lower: "also"},
			},
			1, // all are "addition"
		},
		{
			"filtered_connector_ignored",
			[]Token{
				{Lower: "and", IsFiltered: true},
				{Lower: "cat"},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcConnectorDiversity(tt.tokens)
			if got != tt.want {
				t.Errorf("calcConnectorDiversity() = %d; want %d", got, tt.want)
			}
		})
	}
}

func TestMapConnectorDiversityToScore(t *testing.T) {
	tests := []struct {
		name  string
		count int
		want  float64
	}{
		{"zero", 0, 1.0},
		{"at_2", 2, 1.0},
		{"at_5", 5, 2.0},
		{"at_10", 10, 3.0},
		{"at_15", 15, 4.0},
		{"at_20", 20, 5.0},
		{"at_25", 25, 6.0},
		{"above_max", 30, 6.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapConnectorDiversityToScore(tt.count)
			if !syntaxAlmostEqual(got, tt.want, 0.01) {
				t.Errorf("mapConnectorDiversityToScore(%d) = %.4f; want %.4f", tt.count, got, tt.want)
			}
		})
	}
}

func TestAnalyzeSyntax(t *testing.T) {
	t.Run("empty_input", func(t *testing.T) {
		result := analyzeSyntax(nil, nil)
		if result.Score != 1.0 {
			t.Errorf("empty input: Score = %.4f; want 1.0", result.Score)
		}
	})

	t.Run("simple_sentence", func(t *testing.T) {
		tokens := []Token{
			{Lower: "the"},
			{Lower: "cat"},
			{Lower: "sat"},
		}
		sentences := []string{"The cat sat."}
		result := analyzeSyntax(tokens, sentences)
		if result.AvgSentenceLength != 3.0 {
			t.Errorf("ASL = %.4f; want 3.0", result.AvgSentenceLength)
		}
		if result.Score < 1.0 || result.Score > 6.0 {
			t.Errorf("Score = %.4f; out of [1.0, 6.0]", result.Score)
		}
	})

	t.Run("complex_text", func(t *testing.T) {
		tokens := []Token{
			{Lower: "the"}, {Lower: "book"}, {Lower: "was"},
			{Lower: "written"}, {Lower: "because"}, {Lower: "the"},
			{Lower: "author"}, {Lower: "although"}, {Lower: "tired"},
			{Lower: "however"}, {Lower: "finished"}, {Lower: "and"},
			{Lower: "therefore"}, {Lower: "first"}, {Lower: "overall"},
		}
		sentences := []string{
			"The book was written because the author although tired.",
			"However finished and therefore first overall.",
		}
		result := analyzeSyntax(tokens, sentences)
		if result.Score < 1.0 || result.Score > 6.0 {
			t.Errorf("Score = %.4f; out of [1.0, 6.0]", result.Score)
		}
		if result.AvgSentenceLength != 7.5 {
			t.Errorf("ASL = %.4f; want 7.5", result.AvgSentenceLength)
		}
		// Should detect some subordination (because, although)
		if result.SubordinationIndex < 0.5 {
			t.Errorf("SubordinationIndex = %.4f; expected >= 0.5", result.SubordinationIndex)
		}
		// Should detect connector diversity
		if result.ConnectorDiversity < 3 {
			t.Errorf("ConnectorDiversity = %d; expected >= 3", result.ConnectorDiversity)
		}
	})

	t.Run("filtered_tokens_excluded_from_word_count", func(t *testing.T) {
		tokens := []Token{
			{Lower: "the"},
			{Lower: ".", IsFiltered: true},
			{Lower: "cat"},
		}
		sentences := []string{"The. Cat."}
		result := analyzeSyntax(tokens, sentences)
		if result.AvgSentenceLength != 2.0 {
			t.Errorf("ASL = %.4f; want 2.0 (filtered tokens excluded)", result.AvgSentenceLength)
		}
	})

	t.Run("weights_sum", func(t *testing.T) {
		// Verify that 0.40 + 0.30 + 0.15 + 0.15 = 1.0
		sum := 0.40 + 0.30 + 0.15 + 0.15
		if !syntaxAlmostEqual(sum, 1.0, 0.001) {
			t.Errorf("weights sum = %.4f; want 1.0", sum)
		}
	})
}

func TestCalcPassiveRate_WasEaten(t *testing.T) {
	// "was eaten" is a clear passive (irregular PP).
	// "was walking" is progressive, not passive.
	got := calcPassiveRate([]string{"The cake was eaten by the dog."})
	if !syntaxAlmostEqual(got, 1.0, 0.01) {
		t.Errorf("calcPassiveRate('was eaten') = %.4f; want 1.0", got)
	}

	got2 := calcPassiveRate([]string{"She was walking to the store."})
	if !syntaxAlmostEqual(got2, 0.0, 0.01) {
		t.Errorf("calcPassiveRate('was walking') = %.4f; want 0.0", got2)
	}
}

func TestAnalyzeSyntax_CompositeScore(t *testing.T) {
	// Construct input with known sub-scores to verify weighted composite.
	// 15 non-filtered tokens, 2 sentences → ASL = 7.5
	// Subordinators: because, although → subIdx = 2/2 = 1.0
	// Passive: "was written" in sentence 1 → rate = 1/2 = 0.5
	// Connectors: because(cause), and(addition), moreover(addition),
	//   first(sequence), overall(summary) → 4 distinct categories
	tokens := []Token{
		{Lower: "the"}, {Lower: "book"}, {Lower: "was"},
		{Lower: "written"}, {Lower: "because"}, {Lower: "the"},
		{Lower: "author"}, {Lower: "and"},
		{Lower: "she"}, {Lower: "although"}, {Lower: "tired"},
		{Lower: "moreover"}, {Lower: "ran"}, {Lower: "first"},
		{Lower: "overall"},
	}
	sentences := []string{
		"The book was written because the author and.",
		"She although tired moreover ran first overall.",
	}

	result := analyzeSyntax(tokens, sentences)

	asl := 15.0 / 2.0 // 7.5
	expASL := mapASLToScore(asl)
	expSub := mapSubordinationToScore(1.0)
	expPassive := mapPassiveRateToScore(0.5)
	expConn := mapConnectorDiversityToScore(4)

	wantScore := 0.40*expASL + 0.30*expSub + 0.15*expPassive + 0.15*expConn

	if !syntaxAlmostEqual(result.Score, wantScore, 0.01) {
		t.Errorf("Score = %.4f; want %.4f (0.40×%.2f + 0.30×%.2f + 0.15×%.2f + 0.15×%.2f)",
			result.Score, wantScore, expASL, expSub, expPassive, expConn)
	}
	if !syntaxAlmostEqual(result.AvgSentenceLength, 7.5, 0.01) {
		t.Errorf("ASL = %.4f; want 7.5", result.AvgSentenceLength)
	}
	if !syntaxAlmostEqual(result.SubordinationIndex, 1.0, 0.01) {
		t.Errorf("SubordinationIndex = %.4f; want 1.0", result.SubordinationIndex)
	}
	if result.ConnectorDiversity != 4 {
		t.Errorf("ConnectorDiversity = %d; want 4", result.ConnectorDiversity)
	}
}

func TestExtractWords(t *testing.T) {
	tests := []struct {
		name     string
		sentence string
		want     []string
	}{
		{"simple", "Hello world", []string{"hello", "world"}},
		{"with_punctuation", "Hello, world!", []string{"hello", "world"}},
		{"empty", "", nil},
		{"only_punctuation", ".,!?", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractWords(tt.sentence)
			if len(got) != len(tt.want) {
				t.Fatalf("extractWords(%q) = %v; want %v", tt.sentence, got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("extractWords(%q)[%d] = %q; want %q", tt.sentence, i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestStripNonAlpha(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "hello"},
		{"hello!", "hello"},
		{"(hello)", "hello"},
		{"...word...", "word"},
		{"123", ""},
		{"", ""},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := stripNonAlpha(tt.input)
			if got != tt.want {
				t.Errorf("stripNonAlpha(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}
