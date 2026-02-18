package cefr

import "strings"

// isVowel reports whether c is an English vowel (including y).
func isVowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'y':
		return true
	}
	return false
}

// countSyllables estimates the number of syllables in an English word
// using a vowel-sequence heuristic (~85-90% accuracy).
//
// Rules applied:
//  1. Count consecutive vowel (a, e, i, o, u, y) groups.
//  2. Trailing -es: if preceded by a sibilant (s, x, z, ch, sh) the
//     -es is a syllable (already counted); otherwise subtract 1.
//  3. Trailing -ed: if preceded by t or d the -ed is a syllable
//     (already counted); otherwise subtract 1.
//  4. Trailing -e (not -es/-ed): if preceded by consonant+l (i.e.
//     the word ends in consonant+le) the -le is a syllable (already
//     counted); otherwise subtract 1 for the silent e.
//  5. Result is at least 1.
func countSyllables(word string) int {
	word = strings.ToLower(strings.TrimSpace(word))
	n := len(word)
	if n == 0 {
		return 0
	}

	// Step 1: count consecutive vowel groups.
	count := 0
	inVowel := false
	for i := 0; i < n; i++ {
		if isVowel(word[i]) {
			if !inVowel {
				count++
				inVowel = true
			}
		} else {
			inVowel = false
		}
	}

	// Step 2-4: adjust for special endings.
	switch {
	case n >= 3 && word[n-2:] == "es":
		if !endsWithSibilant(word[:n-2]) {
			count--
		}
	case n >= 3 && word[n-2:] == "ed":
		if word[n-3] != 't' && word[n-3] != 'd' {
			count--
		}
	case n >= 2 && word[n-1] == 'e':
		if n >= 3 && word[n-2] == 'l' && !isVowel(word[n-3]) {
			// consonant + le → syllabic, keep counted
		} else {
			count--
		}
	}

	// Step 5: at least one syllable.
	if count < 1 {
		count = 1
	}
	return count
}

// endsWithSibilant reports whether s ends with a sibilant sound
// (s, x, z, ch, sh) that causes a following -es to form its own
// syllable.
func endsWithSibilant(s string) bool {
	if len(s) == 0 {
		return false
	}
	switch s[len(s)-1] {
	case 's', 'x', 'z':
		return true
	}
	if len(s) >= 2 {
		pair := s[len(s)-2:]
		if pair == "ch" || pair == "sh" {
			return true
		}
	}
	return false
}

// countTotalSyllables returns the sum of syllable counts for all words.
func countTotalSyllables(words []string) int {
	total := 0
	for _, w := range words {
		total += countSyllables(w)
	}
	return total
}
