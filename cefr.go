package cefr

import (
	"errors"
	"strings"
	"unicode"

	"github.com/simp-lee/cefr/data"
)

// Result holds the complete CEFR assessment output for a text.
type Result struct {
	Level         string            // CEFR level label: "A1" through "C2"
	Score         float64           // Continuous score from 1.0 to 6.0
	Confidence    float64           // Confidence of the assessment from 0.0 to 1.0
	Vocab         VocabResult       // Vocabulary analysis details
	Syntax        SyntaxResult      // Syntactic complexity details
	Readability   ReadabilityResult // Readability formula details
	WordCount     int               // Total word count
	SentenceCount int               // Total sentence count
}

// VocabResult holds vocabulary analysis metrics.
type VocabResult struct {
	Score        float64            // Vocabulary sub-score (1.0–6.0)
	Distribution map[string]float64 // CEFR level distribution of content words
	UnknownRatio float64            // Ratio of words not found in any word list
	ContentWords int                // Number of content words analyzed
}

// SyntaxResult holds syntactic complexity metrics.
type SyntaxResult struct {
	Score              float64 // Syntax sub-score (1.0–6.0)
	AvgSentenceLength  float64 // Average words per sentence
	SubordinationIndex float64 // Ratio of subordinate clauses
	PassiveRate        float64 // Ratio of passive voice constructions
	ConnectorDiversity int     // Number of distinct connector types used
}

// ReadabilityResult holds readability formula outputs.
type ReadabilityResult struct {
	Score float64 // Readability sub-score (1.0–6.0)
	FKGL  float64 // Flesch-Kincaid Grade Level
	FRE   float64 // Flesch Reading Ease
	CLI   float64 // Coleman-Liau Index
}

// calcConfidence estimates assessment reliability based on text length and
// inter-segment consistency. Returns a value in [0.0, 1.0].
//
// lengthFactor scales linearly from 0 at 0 words to 1.0 at 500+ words,
// reflecting that longer texts yield more reliable assessments.
//
// consistencyFactor penalises high variance among sampled segments;
// variance >= 1.0 drives confidence to zero.
func calcConfidence(wordCount int, sentenceCount int, sampledVariance float64) float64 {
	lengthFactor := float64(wordCount) / 500.0
	if lengthFactor > 1.0 {
		lengthFactor = 1.0
	}

	consistencyFactor := 1.0 - sampledVariance
	if consistencyFactor < 0.0 {
		consistencyFactor = 0.0
	}

	confidence := lengthFactor * consistencyFactor

	// Clamp to [0.0, 1.0].
	if confidence < 0.0 {
		confidence = 0.0
	}
	if confidence > 1.0 {
		confidence = 1.0
	}
	return confidence
}

// assessText performs CEFR analysis on a single text segment without sampling.
// Confidence is left at 0; the caller is responsible for computing it.
func assessText(text string, cfg config) (Result, error) {
	// Step a: preprocess
	sentences := splitSentences(text)
	tokens := tokenize(text)

	// Step b: input validation – require at least one content word
	hasContent := false
	for _, tok := range tokens {
		if !tok.IsFiltered && !tok.IsStopword && !tok.IsProper {
			hasContent = true
			break
		}
	}
	if !hasContent {
		return Result{}, errors.New("no content words found in text")
	}

	// Step c: vocabulary analysis
	vocab := analyzeVocab(tokens)

	// Step d: syntax analysis
	syntax := analyzeSyntax(tokens, sentences)

	// Step e: syllable and character statistics
	var words []string
	charCount := 0
	wordCount := 0
	for _, tok := range tokens {
		if !tok.IsFiltered {
			words = append(words, tok.Lower)
			charCount += len(tok.Lower)
			wordCount++
		}
	}
	syllableCount := countTotalSyllables(words)

	sentenceCount := len(sentences)
	if sentenceCount == 0 {
		sentenceCount = 1
	}
	if wordCount == 0 {
		wordCount = 1
	}

	// Step f: readability analysis
	readability := analyzeReadability(wordCount, sentenceCount, syllableCount, charCount)

	// Step g: fusion score
	score := vocab.Score*cfg.vocabWeight + syntax.Score*cfg.syntaxWeight + readability.Score*cfg.readabilityWeight

	// Step h: clamp
	score = clampScore(score)

	// Step i: map to level
	level := scoreToLevel(score)

	// Step j/k: assemble Result
	return Result{
		Level:         level,
		Score:         score,
		Confidence:    0, // set by caller
		Vocab:         vocab,
		Syntax:        syntax,
		Readability:   readability,
		WordCount:     wordCount,
		SentenceCount: sentenceCount,
	}, nil
}

// Assess evaluates the CEFR difficulty level of the given English text.
// It returns a Result with the level, score, confidence, and detailed
// sub-scores. Options can customise weights, sampling threshold, etc.
func Assess(text string, opts ...Option) (Result, error) {
	// Step a: validate input
	text = strings.TrimSpace(text)
	if text == "" {
		return Result{}, errors.New("empty text")
	}

	hasLetter := false
	for _, r := range text {
		if unicode.IsLetter(r) && r < 128 {
			hasLetter = true
			break
		}
	}
	if !hasLetter {
		return Result{}, errors.New("no English content")
	}
	if err := data.InitError(); err != nil {
		return Result{}, err
	}

	// Step b: apply options
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	// Step c: quick word count estimate
	wordCount := len(strings.Fields(text))

	// Step d: sampling path
	if shouldSample(wordCount, cfg.samplingThreshold) && !cfg.fullAnalysis {
		segments := sampleText(text, wordCount)
		var results []Result
		for _, seg := range segments {
			r, err := assessText(seg, cfg)
			if err != nil {
				continue // skip segments that error
			}
			results = append(results, r)
		}
		if len(results) == 0 {
			return Result{}, errors.New("all sampled segments failed analysis")
		}

		merged := mergeSampledResults(results, cfg)
		variance := sampledVariance(results)
		merged.Confidence = calcConfidence(merged.WordCount, merged.SentenceCount, variance)
		return merged, nil
	}

	// Step e: full analysis path
	result, err := assessText(text, cfg)
	if err != nil {
		return Result{}, err
	}
	result.Confidence = calcConfidence(result.WordCount, result.SentenceCount, 0)
	return result, nil
}
