package cefr

import (
	"errors"
	"math"
	"strings"
	"sync"
	"testing"
)

func TestCalcConfidence(t *testing.T) {
	tests := []struct {
		name            string
		wordCount       int
		sentenceCount   int
		sampledVariance float64
		want            float64
		wantLessThan    float64 // if > 0, assert result < this
		wantGreaterThan float64 // if > 0, assert result > this
	}{
		{
			name:            "10 words zero variance (FR-104: <0.5)",
			wordCount:       10,
			sentenceCount:   1,
			sampledVariance: 0.0,
			want:            0.02, // 10/500 * 1.0
			wantLessThan:    0.5,
		},
		{
			name:            "1000 words zero variance (FR-104: >0.8)",
			wordCount:       1000,
			sentenceCount:   50,
			sampledVariance: 0.0,
			want:            1.0, // min(1.0, 1000/500) * 1.0
			wantGreaterThan: 0.8,
		},
		{
			name:            "1000 words low variance (FR-104: >0.8)",
			wordCount:       1000,
			sentenceCount:   50,
			sampledVariance: 0.1,
			want:            0.9, // 1.0 * (1.0 - 0.1)
			wantGreaterThan: 0.8,
		},
		{
			name:            "500 words zero variance",
			wordCount:       500,
			sentenceCount:   25,
			sampledVariance: 0.0,
			want:            1.0, // 500/500 * 1.0
		},
		{
			name:            "50 words zero variance",
			wordCount:       50,
			sentenceCount:   3,
			sampledVariance: 0.0,
			want:            0.1, // 50/500 * 1.0
		},
		{
			name:            "250 words half variance",
			wordCount:       250,
			sentenceCount:   15,
			sampledVariance: 0.5,
			want:            0.25, // 0.5 * 0.5
		},
		{
			name:            "zero words",
			wordCount:       0,
			sentenceCount:   0,
			sampledVariance: 0.0,
			want:            0.0,
		},
		{
			name:            "variance >= 1 yields zero",
			wordCount:       1000,
			sentenceCount:   50,
			sampledVariance: 1.0,
			want:            0.0, // 1.0 * max(0, 1.0-1.0) = 0
		},
		{
			name:            "variance > 1 clamped",
			wordCount:       1000,
			sentenceCount:   50,
			sampledVariance: 2.0,
			want:            0.0, // consistencyFactor = max(0, 1-2) = 0
		},
		{
			name:            "negative variance treated as high consistency",
			wordCount:       500,
			sentenceCount:   25,
			sampledVariance: -0.5,
			want:            1.0, // 1.0 * min of clamp: 1.0*(1.0-(-0.5))=1.5 → clamped to 1.0
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcConfidence(tt.wordCount, tt.sentenceCount, tt.sampledVariance)
			if math.Abs(got-tt.want) > 0.001 {
				t.Errorf("calcConfidence(%d, %d, %.2f) = %v; want %v",
					tt.wordCount, tt.sentenceCount, tt.sampledVariance, got, tt.want)
			}
			if tt.wantLessThan > 0 && got >= tt.wantLessThan {
				t.Errorf("calcConfidence(%d, %d, %.2f) = %v; want < %v (FR-104)",
					tt.wordCount, tt.sentenceCount, tt.sampledVariance, got, tt.wantLessThan)
			}
			if tt.wantGreaterThan > 0 && got <= tt.wantGreaterThan {
				t.Errorf("calcConfidence(%d, %d, %.2f) = %v; want > %v (FR-104)",
					tt.wordCount, tt.sentenceCount, tt.sampledVariance, got, tt.wantGreaterThan)
			}
			// Always in [0.0, 1.0]
			if got < 0.0 || got > 1.0 {
				t.Errorf("calcConfidence(%d, %d, %.2f) = %v; out of [0.0, 1.0]",
					tt.wordCount, tt.sentenceCount, tt.sampledVariance, got)
			}
		})
	}
}

func TestAssess_Errors(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		wantErr string
	}{
		{"empty string", "", "empty text"},
		{"only whitespace", "   \n\t  ", "empty text"},
		{"no English content", "12345 67890 !@#$%", "no English content"},
		{"only numbers and symbols", "100 200 300", "no English content"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Assess(tt.text)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("error = %q; want containing %q", err.Error(), tt.wantErr)
			}
		})
	}
}

func TestAssess_SimpleText(t *testing.T) {
	text := "The cat sat on the mat. It was a good day."
	result, err := Assess(text)
	if err != nil {
		t.Fatalf("Assess() error = %v", err)
	}

	// Score must be in [1.0, 6.0]
	if result.Score < 1.0 || result.Score > 6.0 {
		t.Errorf("Score = %v; want in [1.0, 6.0]", result.Score)
	}

	// Level must be a valid CEFR level
	validLevels := map[string]bool{"A1": true, "A2": true, "B1": true, "B2": true, "C1": true, "C2": true}
	if !validLevels[result.Level] {
		t.Errorf("Level = %q; want valid CEFR level", result.Level)
	}

	// Confidence must be in [0.0, 1.0]
	if result.Confidence < 0.0 || result.Confidence > 1.0 {
		t.Errorf("Confidence = %v; want in [0.0, 1.0]", result.Confidence)
	}

	// WordCount and SentenceCount must be positive
	if result.WordCount <= 0 {
		t.Errorf("WordCount = %d; want > 0", result.WordCount)
	}
	if result.SentenceCount <= 0 {
		t.Errorf("SentenceCount = %d; want > 0", result.SentenceCount)
	}

	// Simple text should be low level (A1-B1)
	if result.Score > 4.0 {
		t.Errorf("Simple text Score = %v; expected <= 4.0 for simple text", result.Score)
	}
}

func TestAssess_WithOptions(t *testing.T) {
	text := "The cat sat on the mat. It was a good day."
	r1, err := Assess(text)
	if err != nil {
		t.Fatalf("Assess() error = %v", err)
	}
	r2, err := Assess(text, WithWeights(0.60, 0.20, 0.20))
	if err != nil {
		t.Fatalf("Assess() with weights error = %v", err)
	}
	// Different weights should produce different scores (unless coincidental)
	// At minimum both must be valid
	if r1.Score < 1.0 || r1.Score > 6.0 {
		t.Errorf("r1.Score = %v; want in [1.0, 6.0]", r1.Score)
	}
	if r2.Score < 1.0 || r2.Score > 6.0 {
		t.Errorf("r2.Score = %v; want in [1.0, 6.0]", r2.Score)
	}
}

func TestAssess_ConsistencyScoreLevel(t *testing.T) {
	text := "The cat sat on the mat."
	result, err := Assess(text)
	if err != nil {
		t.Fatalf("Assess() error = %v", err)
	}
	// Level must be consistent with Score
	expectedLevel := scoreToLevel(result.Score)
	if result.Level != expectedLevel {
		t.Errorf("Level = %q inconsistent with Score = %v (expected %q)",
			result.Level, result.Score, expectedLevel)
	}
}

func TestAssessText_NoContentWords(t *testing.T) {
	// Text with only stopwords/filtered tokens should error
	text := "the a an is"
	_, err := assessText(text, defaultConfig())
	if err == nil {
		t.Fatal("expected error for text with no content words, got nil")
	}
}

func TestAssess_SamplingPath(t *testing.T) {
	// Build a text with ~20 words and force sampling with a low threshold.
	text := "The students studied hard for the important examination. They wanted to achieve excellent results in every subject."
	result, err := Assess(text, WithSamplingThreshold(10))
	if err != nil {
		t.Fatalf("Assess() with sampling error = %v", err)
	}

	// Score must be in [1.0, 6.0]
	if result.Score < 1.0 || result.Score > 6.0 {
		t.Errorf("Score = %v; want in [1.0, 6.0]", result.Score)
	}

	// Level must be a valid CEFR level
	validLevels := map[string]bool{"A1": true, "A2": true, "B1": true, "B2": true, "C1": true, "C2": true}
	if !validLevels[result.Level] {
		t.Errorf("Level = %q; want valid CEFR level", result.Level)
	}

	// Confidence must be in [0.0, 1.0]
	if result.Confidence < 0.0 || result.Confidence > 1.0 {
		t.Errorf("Confidence = %v; want in [0.0, 1.0]", result.Confidence)
	}

	// WordCount and SentenceCount must be positive
	if result.WordCount <= 0 {
		t.Errorf("WordCount = %d; want > 0", result.WordCount)
	}
	if result.SentenceCount <= 0 {
		t.Errorf("SentenceCount = %d; want > 0", result.SentenceCount)
	}
}

func TestAssess_SubScoresPopulated(t *testing.T) {
	text := "Although the weather was bad, the students continued studying because they wanted to pass the difficult examination."
	result, err := Assess(text)
	if err != nil {
		t.Fatalf("Assess() error = %v", err)
	}

	// Vocab sub-score in valid range
	if result.Vocab.Score < 1.0 || result.Vocab.Score > 6.0 {
		t.Errorf("Vocab.Score = %v; want in [1.0, 6.0]", result.Vocab.Score)
	}
	// Syntax sub-score in valid range
	if result.Syntax.Score < 1.0 || result.Syntax.Score > 6.0 {
		t.Errorf("Syntax.Score = %v; want in [1.0, 6.0]", result.Syntax.Score)
	}
	// Readability sub-score in valid range
	if result.Readability.Score < 1.0 || result.Readability.Score > 6.0 {
		t.Errorf("Readability.Score = %v; want in [1.0, 6.0]", result.Readability.Score)
	}

	// Final score should be weighted combination (verify closeness)
	cfg := defaultConfig()
	expectedScore := clampScore(
		result.Vocab.Score*cfg.vocabWeight +
			result.Syntax.Score*cfg.syntaxWeight +
			result.Readability.Score*cfg.readabilityWeight,
	)
	if math.Abs(result.Score-expectedScore) > 0.01 {
		t.Errorf("Score = %v; want ~%v (weighted combination)", result.Score, expectedScore)
	}
}

// levelIndex maps a CEFR level string to an integer for ±1 comparison.
func levelIndex(level string) int {
	switch level {
	case "A1":
		return 1
	case "A2":
		return 2
	case "B1":
		return 3
	case "B2":
		return 4
	case "C1":
		return 5
	case "C2":
		return 6
	default:
		return 0
	}
}

func TestAssess_ReferenceTexts(t *testing.T) {
	tests := []struct {
		name          string
		text          string
		expectedLevel string
	}{
		{
			name: "A1 simple daily conversation",
			text: "My name is Tom. I am ten years old. I live in a small house. " +
				"I have a cat and a dog. I like to play in the park. I go to school every day. " +
				"I eat bread and milk for breakfast. My mother is very kind. " +
				"She helps me with my homework. I love my family.",
			expectedLevel: "A1",
		},
		{
			name: "A2 simple description and narrative",
			text: "Last weekend I went to a shopping center with my friends. " +
				"We looked at many different shops and bought some new clothes. " +
				"I found a nice blue jacket that was not too expensive. " +
				"After shopping we had lunch at a small restaurant near the center. " +
				"The food was good and we enjoyed talking together. It was a fun day.",
			expectedLevel: "A2",
		},
		{
			name: "B1 expressing opinions and experiences",
			text: "Technology has changed the way people communicate with each other. " +
				"In the past, people used to write letters which could take days or weeks to arrive. " +
				"Nowadays, we can send messages instantly through email or social media. " +
				"While this is very convenient, some people believe that face-to-face communication " +
				"is becoming less common. It is important to find a balance between using technology " +
				"and maintaining personal relationships.",
			expectedLevel: "B1",
		},
		{
			name: "B2 argumentation and abstraction",
			text: "The increasing prevalence of remote work has fundamentally altered the traditional " +
				"workplace dynamic. Organizations are now compelled to reconsider their operational " +
				"strategies and develop comprehensive policies that address both the advantages and " +
				"challenges of distributed teams. Research suggests that while productivity often " +
				"improves with flexible arrangements, maintaining team cohesion and organizational " +
				"culture becomes significantly more complex in virtual environments.",
			expectedLevel: "B2",
		},
		{
			name: "C1 academic and professional",
			text: "The epistemological implications of artificial intelligence have prompted considerable " +
				"scholarly debate regarding the fundamental nature of cognition and consciousness. " +
				"Contemporary researchers argue that the apparent sophistication of neural network " +
				"architectures, while demonstrating remarkable pattern recognition capabilities, does " +
				"not necessarily constitute genuine understanding in the philosophical sense. This " +
				"distinction between computational prowess and authentic comprehension remains a " +
				"contentious issue within cognitive science discourse.",
			expectedLevel: "C1",
		},
		{
			name: "C2 highly specialized",
			text: "The hermeneutical paradox inherent in poststructuralist literary criticism manifests " +
				"itself through the irreconcilable tension between the purported death of authorial " +
				"intentionality and the inescapable interpretive frameworks that readers inevitably impose " +
				"upon textual artifacts. This epistemological conundrum undermines the very foundations of " +
				"deconstructive methodology, insofar as the practitioner must simultaneously repudiate and " +
				"deploy the logocentric assumptions which undergird Western metaphysical thought.",
			expectedLevel: "C2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Assess(tt.text)
			if err != nil {
				t.Fatalf("Assess() error = %v", err)
			}

			expectedIdx := levelIndex(tt.expectedLevel)
			gotIdx := levelIndex(result.Level)
			if gotIdx == 0 {
				t.Fatalf("Level = %q; not a valid CEFR level", result.Level)
			}

			diff := gotIdx - expectedIdx
			if diff < -1 || diff > 1 {
				t.Errorf("Level = %q (index %d); expected %q (index %d) ±1",
					result.Level, gotIdx, tt.expectedLevel, expectedIdx)
			}

			// Score must be in valid range
			if result.Score < 1.0 || result.Score > 6.0 {
				t.Errorf("Score = %v; want in [1.0, 6.0]", result.Score)
			}

			// Confidence must be in valid range
			if result.Confidence < 0.0 || result.Confidence > 1.0 {
				t.Errorf("Confidence = %v; want in [0.0, 1.0]", result.Confidence)
			}

			t.Logf("text=%q level=%s score=%.2f confidence=%.2f",
				tt.name, result.Level, result.Score, result.Confidence)
		})
	}
}

func TestAssess_Options(t *testing.T) {
	text := "Although the weather was bad, the students continued studying because " +
		"they wanted to pass the difficult examination that was coming next week."

	t.Run("custom weights produce different score", func(t *testing.T) {
		r1, err := Assess(text)
		if err != nil {
			t.Fatalf("Assess() default error = %v", err)
		}
		r2, err := Assess(text, WithWeights(0.60, 0.20, 0.20))
		if err != nil {
			t.Fatalf("Assess() custom weights error = %v", err)
		}

		// Both must be valid
		if r1.Score < 1.0 || r1.Score > 6.0 {
			t.Errorf("r1.Score = %v; want in [1.0, 6.0]", r1.Score)
		}
		if r2.Score < 1.0 || r2.Score > 6.0 {
			t.Errorf("r2.Score = %v; want in [1.0, 6.0]", r2.Score)
		}

		// With different weights on a text where sub-scores differ,
		// we expect different final scores (unless all sub-scores equal).
		if r1.Vocab.Score != r1.Syntax.Score || r1.Syntax.Score != r1.Readability.Score {
			if r1.Score == r2.Score {
				t.Errorf("custom weights should produce different score; both = %v", r1.Score)
			}
		}
	})

	t.Run("WithFullAnalysis returns same result for short text", func(t *testing.T) {
		r1, err := Assess(text)
		if err != nil {
			t.Fatalf("Assess() error = %v", err)
		}
		r2, err := Assess(text, WithFullAnalysis())
		if err != nil {
			t.Fatalf("Assess() WithFullAnalysis error = %v", err)
		}

		// Short text should produce same score regardless of fullAnalysis
		if math.Abs(r1.Score-r2.Score) > 0.001 {
			t.Errorf("WithFullAnalysis() score %v != default score %v for short text",
				r2.Score, r1.Score)
		}
	})

	t.Run("WithSamplingThreshold triggers sampling on longer text", func(t *testing.T) {
		// Build a text with >5 words
		longEnough := "The students studied hard for the important examination every single day."
		r, err := Assess(longEnough, WithSamplingThreshold(5))
		if err != nil {
			t.Fatalf("Assess() with low threshold error = %v", err)
		}

		// Should still produce a valid result after sampling
		if r.Score < 1.0 || r.Score > 6.0 {
			t.Errorf("Score = %v; want in [1.0, 6.0]", r.Score)
		}
		if r.Level == "" {
			t.Error("Level is empty")
		}
	})
}

func TestAssess_EdgeCases(t *testing.T) {
	t.Run("empty text returns error", func(t *testing.T) {
		_, err := Assess("")
		if err == nil {
			t.Fatal("expected error for empty text, got nil")
		}
	})

	t.Run("single word produces low confidence", func(t *testing.T) {
		// Use a content word (non-stopword); "hello" is a stopword.
		result, err := Assess("Beautiful")
		if err != nil {
			t.Fatalf("Assess(\"Beautiful\") error = %v", err)
		}
		if result.Confidence >= 0.5 {
			t.Errorf("Confidence = %v; want < 0.5 for single word", result.Confidence)
		}
		if result.Score < 1.0 || result.Score > 6.0 {
			t.Errorf("Score = %v; want in [1.0, 6.0]", result.Score)
		}
	})

	t.Run("very short text has lower confidence", func(t *testing.T) {
		result, err := Assess("The dog is here now today.")
		if err != nil {
			t.Fatalf("Assess() error = %v", err)
		}
		if result.Confidence >= 0.5 {
			t.Errorf("Confidence = %v; want < 0.5 for very short text", result.Confidence)
		}
	})

	t.Run("pure stopwords returns error", func(t *testing.T) {
		_, err := Assess("the a is in on")
		if err == nil {
			t.Fatal("expected error for pure stopwords, got nil")
		}
	})

	t.Run("all uppercase treated as acronyms", func(t *testing.T) {
		// All-uppercase words (len>1) are classified as acronyms (proper nouns),
		// leaving no content words → error is expected.
		_, err := Assess("THE CAT SAT ON THE MAT")
		if err == nil {
			t.Fatal("expected error for all-uppercase text (words treated as acronyms), got nil")
		}
	})

	t.Run("normal case text is assessed normally", func(t *testing.T) {
		// Regular mixed-case: lowercase content words are not treated as proper nouns.
		result, err := Assess("The cat sat on the mat.")
		if err != nil {
			t.Fatalf("Assess() error = %v", err)
		}
		if result.Score < 1.0 || result.Score > 6.0 {
			t.Errorf("Score = %v; want in [1.0, 6.0]", result.Score)
		}
		if result.Score > 4.0 {
			t.Errorf("Score = %v; expected <= 4.0 for simple text", result.Score)
		}
	})

	t.Run("mixed Chinese-English text filters non-ASCII", func(t *testing.T) {
		result, err := Assess("这是一段 English text 混合的内容 with some words")
		if err != nil {
			t.Fatalf("Assess() error = %v", err)
		}
		if result.Score < 1.0 || result.Score > 6.0 {
			t.Errorf("Score = %v; want in [1.0, 6.0]", result.Score)
		}
	})

	t.Run("pure numbers and punctuation returns error", func(t *testing.T) {
		_, err := Assess("123 456 !!!")
		if err == nil {
			t.Fatal("expected error for pure numbers/punctuation, got nil")
		}
	})

	t.Run("proper nouns do not inflate vocab score", func(t *testing.T) {
		// Text heavy on proper nouns with simple vocabulary
		textWithProper := "Tom and Mary went to London and Paris to visit the British Museum."
		result, err := Assess(textWithProper)
		if err != nil {
			t.Fatalf("Assess() error = %v", err)
		}
		// Even with many proper nouns, simple vocab should keep it low level
		if result.Score > 4.0 {
			t.Errorf("Score = %v; expected <= 4.0 — proper nouns should not inflate vocab",
				result.Score)
		}
		// Verify proper nouns are not counted as high-level unknowns
		if result.Vocab.UnknownRatio > 0.5 {
			t.Errorf("UnknownRatio = %v; expected <= 0.5 — proper nouns should be excluded",
				result.Vocab.UnknownRatio)
		}
	})
}

func TestAssess_Confidence(t *testing.T) {
	t.Run("short text has low confidence", func(t *testing.T) {
		shortText := "The cat sat on the mat. It was good."
		result, err := Assess(shortText)
		if err != nil {
			t.Fatalf("Assess() error = %v", err)
		}
		if result.Confidence >= 0.5 {
			t.Errorf("Confidence = %v; want < 0.5 for ~10 word text", result.Confidence)
		}
	})

	t.Run("long text has higher confidence", func(t *testing.T) {
		// Build a ~200 word text by repeating coherent sentences.
		sentences := []string{
			"Technology has changed the way people communicate with each other.",
			"In the past, people used to write letters which could take days to arrive.",
			"Nowadays, we can send messages instantly through email or social media.",
			"While this is very convenient, some people believe that communication is changing.",
			"It is important to find a balance between technology and personal relationships.",
			"The increasing prevalence of remote work has altered the traditional workplace.",
			"Organizations are compelled to reconsider their operational strategies daily.",
			"Research suggests that productivity often improves with flexible arrangements.",
			"However, maintaining team cohesion becomes more complex in virtual settings.",
			"Companies need to develop comprehensive policies for distributed teams.",
			"Contemporary researchers argue about the nature of cognition and consciousness.",
			"Neural network architectures demonstrate remarkable pattern recognition abilities.",
			"The distinction between computational prowess and understanding remains debated.",
			"Artificial intelligence has prompted considerable scholarly discussion recently.",
			"Education systems worldwide are adapting to incorporate digital learning tools.",
			"Students benefit from access to vast online resources and interactive platforms.",
			"Teachers must balance traditional methods with innovative technological approaches.",
			"The global economy continues to evolve in response to technological advancement.",
			"International collaboration has become essential for addressing complex challenges.",
			"Environmental policies require careful consideration of economic and social factors.",
		}
		longText := strings.Join(sentences, " ")
		result, err := Assess(longText)
		if err != nil {
			t.Fatalf("Assess() error = %v", err)
		}

		if result.Confidence <= 0.3 {
			t.Errorf("Confidence = %v; want > 0.3 for ~200 word text", result.Confidence)
		}

		// Long text confidence should be notably higher than short text
		shortResult, _ := Assess("The cat sat on the mat.")
		if result.Confidence <= shortResult.Confidence {
			t.Errorf("long text confidence (%v) should exceed short text confidence (%v)",
				result.Confidence, shortResult.Confidence)
		}
	})
}

func TestAssess_ConcurrentSafety(t *testing.T) {
	texts := []string{
		"My name is Tom. I am ten years old. I live in a small house.",
		"Technology has changed the way people communicate with each other.",
		"The increasing prevalence of remote work has fundamentally altered workplace dynamics.",
		"The epistemological implications of artificial intelligence have prompted scholarly debate.",
	}

	var wg sync.WaitGroup
	errs := make(chan error, len(texts)*10)

	for i := 0; i < 10; i++ {
		for _, text := range texts {
			wg.Add(1)
			go func(txt string) {
				defer wg.Done()
				result, err := Assess(txt)
				if err != nil {
					errs <- err
					return
				}
				if result.Score < 1.0 || result.Score > 6.0 {
					errs <- errors.New("score out of range in concurrent call")
					return
				}
			}(text)
		}
	}

	wg.Wait()
	close(errs)

	for err := range errs {
		t.Errorf("concurrent Assess() error: %v", err)
	}
}
