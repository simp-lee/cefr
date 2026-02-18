package cefr

import (
	"sort"

	"github.com/simp-lee/cefr/data"
)

// lookupWordLevel returns the CEFR integer level (1–6) for a word by checking
// Oxford 5000, AWL, and NGSL in that order. It first tries the original form,
// then the lemmatized form. Returns (0, false) if not found in any list.
func lookupWordLevel(word string) (int, bool) {
	if word == "" {
		return 0, false
	}

	// Try original word first
	if level, ok := lookupInLists(word); ok {
		return level, true
	}

	// Try lemmatized form
	lemma := lemmatize(word)
	if lemma != word {
		if level, ok := lookupInLists(lemma); ok {
			return level, true
		}
	}

	return 0, false
}

// lookupInLists checks a single word form against Oxford 5000, AWL, then NGSL.
func lookupInLists(word string) (int, bool) {
	if level, ok := data.OxfordLevel(word); ok {
		return level, true
	}

	if level, ok := data.AWLLevel(word); ok {
		return level, true
	}

	if level, ok := data.NGSLLevel(word); ok {
		return level, true
	}

	return 0, false
}

// analyzeVocab computes vocabulary-level metrics from a list of tokens.
// It filters out stopwords, proper nouns, and filtered tokens, deduplicates
// by lowercase form, and computes a score using the P80 percentile method.
func analyzeVocab(tokens []Token) VocabResult {
	if len(tokens) == 0 {
		return VocabResult{
			Score:        1.0,
			Distribution: map[string]float64{},
		}
	}

	// Collect unique content words (deduplicate by Lower form)
	seen := make(map[string]bool)
	var levels []int
	unknownCount := 0

	for _, tok := range tokens {
		if tok.IsStopword || tok.IsProper || tok.IsFiltered {
			continue
		}
		if seen[tok.Lower] {
			continue
		}
		seen[tok.Lower] = true

		level, found := lookupWordLevel(tok.Lower)
		if !found {
			unknownCount++
			// Unknown words assumed C1 (=5) for scoring (FR-202)
			levels = append(levels, int(C1))
		} else {
			levels = append(levels, level)
		}
	}

	contentWords := len(levels)
	if contentWords == 0 {
		return VocabResult{
			Score:        1.0,
			Distribution: map[string]float64{},
		}
	}

	// Calculate unknown ratio
	unknownRatio := float64(unknownCount) / float64(contentWords)

	// Calculate distribution
	distribution := computeDistribution(levels, unknownCount, contentWords)

	// Calculate score using P80 percentile (FR-209)
	sort.Ints(levels)
	var score float64
	if contentWords < 10 {
		// Too few words: use average
		sum := 0
		for _, l := range levels {
			sum += l
		}
		score = float64(sum) / float64(contentWords)
	} else {
		// P80: 80th percentile
		idx := int(float64(contentWords-1) * 0.8)
		score = float64(levels[idx])
	}

	// Unknown word correction: heavy unknown ratio pushes score up
	if unknownRatio > 0.3 {
		score += (unknownRatio - 0.3) * 2.0
	}

	// Clamp to [1.0, 6.0]
	score = clampScore(score)

	return VocabResult{
		Score:        score,
		Distribution: distribution,
		UnknownRatio: unknownRatio,
		ContentWords: contentWords,
	}
}

// computeDistribution calculates the percentage of content words at each CEFR
// level plus an "Unknown" category.
func computeDistribution(levels []int, unknownCount, total int) map[string]float64 {
	counts := make(map[string]int)
	for _, l := range levels {
		label := Level(l).String()
		counts[label]++
	}

	// Unknown words were added to levels as C1, so we need to adjust:
	// subtract unknownCount from C1 count and add to Unknown
	if unknownCount > 0 {
		counts[C1.String()] -= unknownCount
		if counts[C1.String()] <= 0 {
			delete(counts, C1.String())
		}
	}

	dist := make(map[string]float64)
	for _, label := range []string{"A1", "A2", "B1", "B2", "C1", "C2"} {
		if c, ok := counts[label]; ok && c > 0 {
			dist[label] = float64(c) / float64(total)
		}
	}
	if unknownCount > 0 {
		dist["Unknown"] = float64(unknownCount) / float64(total)
	}

	return dist
}
