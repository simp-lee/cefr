package cefr

import (
	"strings"

	"github.com/simp-lee/cefr/data"
)

// lemmatize attempts to reduce a lowercase English word to its base (lemma) form.
// It first checks the irregular variant table, then tries rule-based suffix stripping.
// If the stripped form is found in the vocabulary, it is returned; otherwise, the
// original word is returned unchanged.
func lemmatize(word string) string {
	if word == "" {
		return word
	}

	// Step 1: Check irregular forms (went→go, better→good, children→child)
	if lemma, ok := data.IrregularLemma(word); ok {
		return lemma
	}

	// Step 2: If already in vocab, return as-is
	if isInVocab(word) {
		return word
	}

	// Step 3: Try rule-based suffix stripping (in priority order)
	// Each rule tries to strip a suffix and validate the result against the vocabulary.
	if result, ok := tryStripIng(word); ok {
		return result
	}
	if result, ok := tryStripEd(word); ok {
		return result
	}
	if result, ok := tryStripEs(word); ok {
		return result
	}
	if result, ok := tryStripS(word); ok {
		return result
	}
	if result, ok := tryStripEr(word); ok {
		return result
	}
	if result, ok := tryStripEst(word); ok {
		return result
	}
	if result, ok := tryStripLy(word); ok {
		return result
	}

	return word
}

// tryStripIng handles -ing suffix:
//   - doubled consonant: running→run
//   - restore e: making→make
//   - direct strip: walking→walk
func tryStripIng(word string) (string, bool) {
	if !strings.HasSuffix(word, "ing") || len(word) < 5 {
		return "", false
	}
	stem := word[:len(word)-3]

	// Doubled consonant: running→runn→run
	if isDoubledConsonant(stem) {
		candidate := stem[:len(stem)-1]
		if isInVocab(candidate) {
			return candidate, true
		}
	}

	// Restore e: making→mak→make
	candidate := stem + "e"
	if isInVocab(candidate) {
		return candidate, true
	}

	// Direct strip: walking→walk
	if isInVocab(stem) {
		return stem, true
	}

	return "", false
}

// tryStripEd handles -ed suffix:
//   - ied→y: studied→study
//   - doubled consonant: stopped→stop
//   - restore e: hoped→hope
//   - direct strip: walked→walk
func tryStripEd(word string) (string, bool) {
	if !strings.HasSuffix(word, "ed") || len(word) < 4 {
		return "", false
	}

	// ied→y: studied→study
	if strings.HasSuffix(word, "ied") && len(word) >= 5 {
		candidate := word[:len(word)-3] + "y"
		if isInVocab(candidate) {
			return candidate, true
		}
	}

	stem := word[:len(word)-2]

	// Doubled consonant: stopped→stopp→stop
	if isDoubledConsonant(stem) {
		candidate := stem[:len(stem)-1]
		if isInVocab(candidate) {
			return candidate, true
		}
	}

	// Restore e: hoped→hop→hope
	candidate := stem + "e"
	if isInVocab(candidate) {
		return candidate, true
	}

	// Direct strip: walked→walk
	if isInVocab(stem) {
		return stem, true
	}

	return "", false
}

// tryStripEs handles -es suffix:
//   - ies→y: stories→story
//   - xes/shes/ches/sses → strip es: boxes→box, dishes→dish, watches→watch
//   - fallback: strip s
func tryStripEs(word string) (string, bool) {
	if !strings.HasSuffix(word, "es") || len(word) < 4 {
		return "", false
	}

	// ies→y: stories→story
	if strings.HasSuffix(word, "ies") && len(word) >= 5 {
		candidate := word[:len(word)-3] + "y"
		if isInVocab(candidate) {
			return candidate, true
		}
	}

	// xes/shes/ches/sses → strip es: boxes→box, dishes→dish, watches→watch
	stem := word[:len(word)-2]
	if strings.HasSuffix(stem, "x") || strings.HasSuffix(stem, "sh") ||
		strings.HasSuffix(stem, "ch") || strings.HasSuffix(stem, "ss") ||
		strings.HasSuffix(stem, "z") {
		if isInVocab(stem) {
			return stem, true
		}
	}

	// Fallback: strip s (not es)
	candidate := word[:len(word)-1]
	if isInVocab(candidate) {
		return candidate, true
	}

	return "", false
}

// tryStripS handles -s suffix (simple plural / 3rd person singular):
//   - direct strip: makes→make
func tryStripS(word string) (string, bool) {
	if !strings.HasSuffix(word, "s") || len(word) < 3 {
		return "", false
	}
	// Skip words ending in "ss" (e.g., "miss", "boss")
	if strings.HasSuffix(word, "ss") {
		return "", false
	}
	candidate := word[:len(word)-1]
	if isInVocab(candidate) {
		return candidate, true
	}
	return "", false
}

// tryStripEr handles -er suffix (comparative):
//   - ier→y: happier→happy
//   - doubled consonant: bigger→big
//   - direct strip: walker→walk
func tryStripEr(word string) (string, bool) {
	if !strings.HasSuffix(word, "er") || len(word) < 4 {
		return "", false
	}

	// ier→y: happier→happy
	if strings.HasSuffix(word, "ier") && len(word) >= 5 {
		candidate := word[:len(word)-3] + "y"
		if isInVocab(candidate) {
			return candidate, true
		}
	}

	stem := word[:len(word)-2]

	// Doubled consonant: bigger→bigg→big
	if isDoubledConsonant(stem) {
		candidate := stem[:len(stem)-1]
		if isInVocab(candidate) {
			return candidate, true
		}
	}

	// Restore e: nicer→nic→nice
	candidate := stem + "e"
	if isInVocab(candidate) {
		return candidate, true
	}

	// Direct strip: walker→walk
	if isInVocab(stem) {
		return stem, true
	}

	return "", false
}

// tryStripEst handles -est suffix (superlative):
//   - iest→y: happiest→happy
//   - doubled consonant: biggest→big
//   - direct strip
func tryStripEst(word string) (string, bool) {
	if !strings.HasSuffix(word, "est") || len(word) < 5 {
		return "", false
	}

	// iest→y: happiest→happy
	if strings.HasSuffix(word, "iest") && len(word) >= 6 {
		candidate := word[:len(word)-4] + "y"
		if isInVocab(candidate) {
			return candidate, true
		}
	}

	stem := word[:len(word)-3]

	// Doubled consonant: biggest→bigg→big
	if isDoubledConsonant(stem) {
		candidate := stem[:len(stem)-1]
		if isInVocab(candidate) {
			return candidate, true
		}
	}

	// Restore e: nicest→nic→nice
	candidate := stem + "e"
	if isInVocab(candidate) {
		return candidate, true
	}

	// Direct strip
	if isInVocab(stem) {
		return stem, true
	}

	return "", false
}

// tryStripLy handles -ly suffix (adverb):
//   - ily→y: happily→happy
//   - direct strip: quickly→quick
func tryStripLy(word string) (string, bool) {
	if !strings.HasSuffix(word, "ly") || len(word) < 4 {
		return "", false
	}

	// ily→y: happily→happy
	if strings.HasSuffix(word, "ily") && len(word) >= 5 {
		candidate := word[:len(word)-3] + "y"
		if isInVocab(candidate) {
			return candidate, true
		}
	}

	// Direct strip: quickly→quick
	stem := word[:len(word)-2]
	if isInVocab(stem) {
		return stem, true
	}

	return "", false
}

// isDoubledConsonant checks if a word ends with a doubled consonant
// (e.g., "runn" → true, "walk" → false).
func isDoubledConsonant(word string) bool {
	if len(word) < 2 {
		return false
	}
	last := word[len(word)-1]
	prev := word[len(word)-2]
	return last == prev && isConsonant(last)
}

// isConsonant returns true if the byte is an English consonant (not a vowel).
func isConsonant(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'y':
		return false
	default:
		return b >= 'a' && b <= 'z'
	}
}

// isInVocab checks if a word exists in the Oxford 5000 or NGSL word lists.
func isInVocab(word string) bool {
	if _, ok := data.OxfordLevel(word); ok {
		return true
	}
	_, ok := data.NGSLLevel(word)
	return ok
}
