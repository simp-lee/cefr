package cefr

import (
	"strings"
	"unicode"

	"github.com/simp-lee/cefr/data"
)

// subordinators is the set of subordinating conjunctions and relative pronouns
// used to compute the subordination index.
var subordinators = map[string]bool{
	"because": true, "although": true, "which": true, "who": true,
	"if": true, "when": true, "while": true, "unless": true,
	"though": true, "since": true, "that": true, "where": true,
	"whose": true, "whom": true, "whereas": true, "whereby": true,
	"whenever": true, "wherever": true, "however": true, "whatever": true,
	"whichever": true, "whoever": true, "after": true, "before": true,
	"until": true, "as": true, "once": true, "provided": true,
	"supposing": true, "lest": true,
}

// beVerbs is the set of be-verb forms used for passive voice detection.
var beVerbs = map[string]bool{
	"is": true, "am": true, "are": true, "was": true,
	"were": true, "been": true, "being": true, "be": true,
}

// connectorCategory maps connector words to their semantic category.
// Multi-word connectors are represented by their distinctive content word.
var connectorCategory = map[string]string{
	// addition
	"and": "addition", "also": "addition", "furthermore": "addition",
	"moreover": "addition", "besides": "addition", "additionally": "addition",
	"addition": "addition", "too": "addition", "likewise": "addition",
	"similarly": "addition", "equally": "addition",
	// contrast
	"but": "contrast", "however": "contrast", "nevertheless": "contrast",
	"nonetheless": "contrast", "yet": "contrast", "still": "contrast",
	"conversely": "contrast", "despite": "contrast", "contrary": "contrast",
	// cause
	"because": "cause", "therefore": "cause", "consequently": "cause",
	"thus": "cause", "hence": "cause", "so": "cause", "accordingly": "cause",
	// sequence
	"first": "sequence", "second": "sequence", "third": "sequence",
	"finally": "sequence", "then": "sequence", "next": "sequence",
	"subsequently": "sequence", "previously": "sequence",
	"meanwhile": "sequence", "afterwards": "sequence",
	"initially": "sequence", "ultimately": "sequence", "lastly": "sequence",
	// summary
	"conclusion": "summary", "summary": "summary",
	"overall": "summary", "briefly": "summary", "conclude": "summary",
}

// isRegularPastParticiple reports whether word looks like a regular past
// participle (ends in -ed and does not end in -ing).
func isRegularPastParticiple(word string) bool {
	return strings.HasSuffix(word, "ed") && !strings.HasSuffix(word, "ing")
}

// calcAvgSentenceLength returns the average number of words per sentence.
func calcAvgSentenceLength(wordCount, sentenceCount int) float64 {
	if sentenceCount == 0 {
		return 0
	}
	return float64(wordCount) / float64(sentenceCount)
}

// aslAnchors maps average sentence length to CEFR score (1.0–6.0).
var aslAnchors = [][2]float64{
	{6, 1.0},
	{9, 2.0},
	{13, 3.0},
	{18, 4.0},
	{23, 5.0},
	{28, 6.0},
}

// mapASLToScore maps an average sentence length to a CEFR score via
// piecewise linear interpolation.
func mapASLToScore(asl float64) float64 {
	return interpolateScore(aslAnchors, asl)
}

// calcSubordinationIndex returns the ratio of subordinating conjunctions
// and relative pronouns to the number of sentences.
func calcSubordinationIndex(tokens []Token, sentenceCount int) float64 {
	if sentenceCount == 0 {
		return 0
	}
	count := 0
	for _, t := range tokens {
		if !t.IsFiltered && subordinators[t.Lower] {
			count++
		}
	}
	return float64(count) / float64(sentenceCount)
}

// subordinationAnchors maps subordination index to CEFR score (1.0–6.0).
var subordinationAnchors = [][2]float64{
	{0, 1.0},
	{0.3, 2.0},
	{0.6, 3.0},
	{1.0, 4.0},
	{1.5, 5.0},
	{2.0, 6.0},
}

// mapSubordinationToScore maps a subordination index to a CEFR score via
// piecewise linear interpolation.
func mapSubordinationToScore(index float64) float64 {
	return interpolateScore(subordinationAnchors, index)
}

// calcPassiveRate estimates the ratio of passive voice constructions
// (be-verb + past participle bigrams) per sentence.
func calcPassiveRate(sentences []string) float64 {
	if len(sentences) == 0 {
		return 0
	}
	irregularPP := data.LoadIrregularPastParticiples()
	passiveCount := 0
	for _, sent := range sentences {
		words := extractWords(sent)
		for i := 0; i < len(words)-1; i++ {
			if beVerbs[words[i]] && isPastParticiple(words[i+1], irregularPP) {
				passiveCount++
			}
		}
	}
	return float64(passiveCount) / float64(len(sentences))
}

// extractWords splits a sentence into lowercase alphabetic words,
// stripping leading and trailing punctuation from each token.
func extractWords(sentence string) []string {
	fields := strings.Fields(sentence)
	words := make([]string, 0, len(fields))
	for _, f := range fields {
		w := stripNonAlpha(f)
		if w != "" {
			words = append(words, strings.ToLower(w))
		}
	}
	return words
}

// stripNonAlpha removes leading and trailing non-letter characters from s.
func stripNonAlpha(s string) string {
	runes := []rune(s)
	start := 0
	for start < len(runes) && !unicode.IsLetter(runes[start]) {
		start++
	}
	end := len(runes)
	for end > start && !unicode.IsLetter(runes[end-1]) {
		end--
	}
	return string(runes[start:end])
}

// isPastParticiple checks whether word is a past participle form,
// first checking the irregular past participles table, then applying
// the regular -ed suffix rule.
func isPastParticiple(word string, irregularPP map[string]bool) bool {
	if irregularPP[word] {
		return true
	}
	return isRegularPastParticiple(word)
}

// passiveAnchors maps passive rate to CEFR score (1.0–6.0).
var passiveAnchors = [][2]float64{
	{0, 1.0},
	{0.05, 2.0},
	{0.10, 3.0},
	{0.20, 4.0},
	{0.30, 5.0},
	{0.40, 6.0},
}

// mapPassiveRateToScore maps a passive voice rate to a CEFR score via
// piecewise linear interpolation.
func mapPassiveRateToScore(rate float64) float64 {
	return interpolateScore(passiveAnchors, rate)
}

// calcConnectorDiversity counts the number of distinct connector categories
// (semantic types) found in the tokens.
func calcConnectorDiversity(tokens []Token) int {
	seen := make(map[string]bool)
	for _, t := range tokens {
		if !t.IsFiltered {
			if category := connectorCategory[t.Lower]; category != "" {
				seen[category] = true
			}
		}
	}
	return len(seen)
}

// connectorAnchors maps connector diversity count to CEFR score (1.0–6.0).
var connectorAnchors = [][2]float64{
	{2, 1.0},
	{5, 2.0},
	{10, 3.0},
	{15, 4.0},
	{20, 5.0},
	{25, 6.0},
}

// mapConnectorDiversityToScore maps the count of distinct connector words
// to a CEFR score via piecewise linear interpolation.
func mapConnectorDiversityToScore(count int) float64 {
	return interpolateScore(connectorAnchors, float64(count))
}

// analyzeSyntax computes syntactic complexity metrics and a combined CEFR score.
//
// Weights: 40% ASL + 30% subordination + 15% passive + 15% connector.
// Returns a default SyntaxResult with Score 1.0 when input is empty.
func analyzeSyntax(tokens []Token, sentences []string) SyntaxResult {
	sentenceCount := len(sentences)

	// Count non-filtered tokens as words.
	wordCount := 0
	for _, t := range tokens {
		if !t.IsFiltered {
			wordCount++
		}
	}

	if wordCount == 0 || sentenceCount == 0 {
		return SyntaxResult{Score: 1.0}
	}

	asl := calcAvgSentenceLength(wordCount, sentenceCount)
	subIdx := calcSubordinationIndex(tokens, sentenceCount)
	passive := calcPassiveRate(sentences)
	connDiv := calcConnectorDiversity(tokens)

	aslScore := mapASLToScore(asl)
	subScore := mapSubordinationToScore(subIdx)
	passiveScore := mapPassiveRateToScore(passive)
	connScore := mapConnectorDiversityToScore(connDiv)

	score := 0.40*aslScore + 0.30*subScore + 0.15*passiveScore + 0.15*connScore

	return SyntaxResult{
		Score:              score,
		AvgSentenceLength:  asl,
		SubordinationIndex: subIdx,
		PassiveRate:        passive,
		ConnectorDiversity: connDiv,
	}
}
