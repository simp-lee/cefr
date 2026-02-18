package cefr

import (
	"math"
	"testing"
)

func TestLookupWordLevel(t *testing.T) {
	tests := []struct {
		name      string
		word      string
		wantLevel int
		wantFound bool
	}{
		// Oxford 5000 words
		{"oxford_a1_word", "book", 1, true},
		{"oxford_a1_important", "important", 1, true},
		{"oxford_a2_environment", "environment", 2, true},
		{"oxford_a2_comprehensive", "comprehensive", 2, true},
		{"oxford_c1_arbitrary", "arbitrary", 5, true},
		{"oxford_b1_word", "elaborate", 3, true},

		// AWL words (B2-C1 range, levels 4-5)
		{"awl_word", "analyse", 4, true},

		// NGSL words (common words)
		{"ngsl_common_word", "the", 1, true},

		// Inflected form that lemmatizes to a known word
		{"inflected_running", "running", 1, true},
		{"inflected_studied", "studied", 1, true},

		// Unknown word
		{"unknown_word", "xylophonic", 0, false},
		{"empty_string", "", 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLevel, gotFound := lookupWordLevel(tt.word)
			if gotFound != tt.wantFound {
				t.Errorf("lookupWordLevel(%q) found=%v, want %v", tt.word, gotFound, tt.wantFound)
			}
			if gotLevel != tt.wantLevel {
				t.Errorf("lookupWordLevel(%q) level=%d, want %d", tt.word, gotLevel, tt.wantLevel)
			}
		})
	}
}

func TestAnalyzeVocab_Basic(t *testing.T) {
	// Create tokens simulating "The cat sat on a mat" where "the","on","a" are stopwords.
	tokens := []Token{
		{Original: "The", Lower: "the", IsStopword: true},
		{Original: "cat", Lower: "cat"},
		{Original: "sat", Lower: "sat"},
		{Original: "on", Lower: "on", IsStopword: true},
		{Original: "a", Lower: "a", IsStopword: true},
		{Original: "mat", Lower: "mat"},
	}

	result := analyzeVocab(tokens)

	// Should have 3 content words: cat, sat, mat
	if result.ContentWords != 3 {
		t.Errorf("ContentWords=%d, want 3", result.ContentWords)
	}

	// Score should be in valid range
	if result.Score < 1.0 || result.Score > 6.0 {
		t.Errorf("Score=%f, want in [1.0, 6.0]", result.Score)
	}

	// Distribution should be non-nil
	if result.Distribution == nil {
		t.Fatal("Distribution is nil")
	}
}

func TestAnalyzeVocab_FiltersProperAndFiltered(t *testing.T) {
	tokens := []Token{
		{Original: "John", Lower: "john", IsProper: true},
		{Original: "runs", Lower: "runs"},
		{Original: "123", Lower: "123", IsFiltered: true},
		{Original: "fast", Lower: "fast"},
	}

	result := analyzeVocab(tokens)

	// Should only count "runs" and "fast" as content words
	if result.ContentWords != 2 {
		t.Errorf("ContentWords=%d, want 2", result.ContentWords)
	}
}

func TestAnalyzeVocab_Deduplication(t *testing.T) {
	// Same word appearing multiple times should be counted once
	tokens := []Token{
		{Original: "run", Lower: "run"},
		{Original: "run", Lower: "run"},
		{Original: "run", Lower: "run"},
		{Original: "fast", Lower: "fast"},
	}

	result := analyzeVocab(tokens)

	// Should have 2 unique content words: run, fast
	if result.ContentWords != 2 {
		t.Errorf("ContentWords=%d, want 2", result.ContentWords)
	}
}

func TestAnalyzeVocab_UnknownRatio(t *testing.T) {
	// All unknown words should yield unknownRatio = 1.0
	tokens := []Token{
		{Original: "xyzfoo", Lower: "xyzfoo"},
		{Original: "xyzbar", Lower: "xyzbar"},
		{Original: "xyzbaz", Lower: "xyzbaz"},
	}

	result := analyzeVocab(tokens)

	if math.Abs(result.UnknownRatio-1.0) > 0.01 {
		t.Errorf("UnknownRatio=%f, want ~1.0", result.UnknownRatio)
	}
}

func TestAnalyzeVocab_EmptyTokens(t *testing.T) {
	result := analyzeVocab(nil)

	if result.ContentWords != 0 {
		t.Errorf("ContentWords=%d, want 0", result.ContentWords)
	}
	if result.Score != 1.0 {
		t.Errorf("Score=%f, want 1.0 for empty input", result.Score)
	}
}

func TestAnalyzeVocab_FewContentWords_UsesAverage(t *testing.T) {
	// Fewer than 10 content words should use average instead of P80
	tokens := []Token{
		{Original: "book", Lower: "book"},   // A1 → 1
		{Original: "water", Lower: "water"}, // A1 → 1
		{Original: "house", Lower: "house"}, // A1 → 1
	}

	result := analyzeVocab(tokens)

	// All are level 1 (A1), average should be ~1.0
	if result.Score < 1.0 || result.Score > 2.0 {
		t.Errorf("Score=%f, expected close to 1.0 for simple A1 words", result.Score)
	}
}

func TestAnalyzeVocab_UnknownCorrectionApplied(t *testing.T) {
	// Create many unique unknown words (>30% unknown) to trigger correction
	var tokens []Token
	// 3 known words
	tokens = append(tokens, Token{Original: "book", Lower: "book"})
	tokens = append(tokens, Token{Original: "water", Lower: "water"})
	tokens = append(tokens, Token{Original: "house", Lower: "house"})

	// 7 unknown words → 70% unknown ratio
	for i := range 7 {
		w := "xyzunknown" + string(rune('a'+i))
		tokens = append(tokens, Token{Original: w, Lower: w})
	}

	result := analyzeVocab(tokens)

	// unknownRatio = 0.7, correction = (0.7 - 0.3) * 2.0 = 0.8
	// Score should be higher than without correction
	if result.UnknownRatio < 0.5 {
		t.Errorf("UnknownRatio=%f, expected >0.5", result.UnknownRatio)
	}
	// Score should be clamped to [1.0, 6.0]
	if result.Score < 1.0 || result.Score > 6.0 {
		t.Errorf("Score=%f, want in [1.0, 6.0]", result.Score)
	}
}

func TestAnalyzeVocab_DistributionSumsToOne(t *testing.T) {
	tokens := []Token{
		{Original: "book", Lower: "book"},
		{Original: "water", Lower: "water"},
		{Original: "elaborate", Lower: "elaborate"},
		{Original: "xyzfoo", Lower: "xyzfoo"},
	}

	result := analyzeVocab(tokens)

	var total float64
	for _, v := range result.Distribution {
		total += v
	}
	if math.Abs(total-1.0) > 0.01 {
		t.Errorf("Distribution total=%f, want ~1.0", total)
	}
}

func TestAnalyzeVocab_DistributionPercentages(t *testing.T) {
	// 4 content words: 2×A1(book,water), 1×B1(elaborate), 1 unknown(xyzfoo)
	tokens := []Token{
		{Original: "book", Lower: "book"},           // A1
		{Original: "water", Lower: "water"},         // A1
		{Original: "elaborate", Lower: "elaborate"}, // B1
		{Original: "xyzfoo", Lower: "xyzfoo"},       // Unknown
	}

	result := analyzeVocab(tokens)

	// A1 = 2/4 = 0.50
	if math.Abs(result.Distribution["A1"]-0.50) > 0.01 {
		t.Errorf("Distribution[A1]=%f, want 0.50", result.Distribution["A1"])
	}
	// B1 = 1/4 = 0.25
	if math.Abs(result.Distribution["B1"]-0.25) > 0.01 {
		t.Errorf("Distribution[B1]=%f, want 0.25", result.Distribution["B1"])
	}
	// Unknown = 1/4 = 0.25
	if math.Abs(result.Distribution["Unknown"]-0.25) > 0.01 {
		t.Errorf("Distribution[Unknown]=%f, want 0.25", result.Distribution["Unknown"])
	}
}

func TestAnalyzeVocab_P80Calculation(t *testing.T) {
	// Build 10 content words with known levels to exercise P80.
	// Levels: book=1,water=1,house=1,good=1,name=1,important=1,environment=2,develop=2,elaborate=3,arbitrary=5
	// Sorted: [1,1,1,1,1,1,2,2,3,5]
	// P80 index = int(9 * 0.8) = 7  → levels[7] = 2
	tokens := []Token{
		{Original: "book", Lower: "book"},               // A1=1
		{Original: "water", Lower: "water"},             // A1=1
		{Original: "house", Lower: "house"},             // A1=1
		{Original: "good", Lower: "good"},               // A1=1
		{Original: "name", Lower: "name"},               // A1=1
		{Original: "important", Lower: "important"},     // A1=1
		{Original: "environment", Lower: "environment"}, // A2=2
		{Original: "develop", Lower: "develop"},         // A2=2
		{Original: "elaborate", Lower: "elaborate"},     // B1=3
		{Original: "arbitrary", Lower: "arbitrary"},     // C1=5
	}

	result := analyzeVocab(tokens)

	if result.ContentWords != 10 {
		t.Fatalf("ContentWords=%d, want 10", result.ContentWords)
	}
	// With 10 words, P80 is used (not average).
	// No unknown words, so no correction applied.
	if result.UnknownRatio > 0.01 {
		t.Errorf("UnknownRatio=%f, expected 0", result.UnknownRatio)
	}
	// P80 = levels[7] = 2.0
	if math.Abs(result.Score-2.0) > 0.01 {
		t.Errorf("Score=%f, want 2.0 (P80 of sorted levels)", result.Score)
	}
}

func TestAnalyzeVocab_AllStopwords(t *testing.T) {
	tokens := []Token{
		{Original: "The", Lower: "the", IsStopword: true},
		{Original: "a", Lower: "a", IsStopword: true},
		{Original: "is", Lower: "is", IsStopword: true},
		{Original: "on", Lower: "on", IsStopword: true},
	}

	result := analyzeVocab(tokens)

	if result.ContentWords != 0 {
		t.Errorf("ContentWords=%d, want 0", result.ContentWords)
	}
	if result.Score != 1.0 {
		t.Errorf("Score=%f, want 1.0", result.Score)
	}
	if result.UnknownRatio != 0 {
		t.Errorf("UnknownRatio=%f, want 0", result.UnknownRatio)
	}
	if len(result.Distribution) != 0 {
		t.Errorf("Distribution=%v, want empty map", result.Distribution)
	}
}
