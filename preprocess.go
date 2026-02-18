package cefr

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/simp-lee/cefr/data"
)

// Token represents a processed word from the input text.
type Token struct {
	Original   string // Original text (case-preserved)
	Lower      string // Lowercase form (for dictionary lookup)
	IsFirst    bool   // Whether this is a sentence-initial word
	IsProper   bool   // Whether this is a proper noun (skip for scoring)
	IsStopword bool   // Whether this is a stopword
	IsFiltered bool   // Whether this is filtered (number/punctuation/non-ASCII)
}

// tokenize splits raw English text into a structured list of Tokens
// with metadata for vocabulary and syntax analysis.
func tokenize(text string) []Token {
	// Step 1: Normalize Unicode characters and remove non-ASCII
	text = normalizeAndFilterASCII(text)

	// Step 2: Split by whitespace
	rawTokens := strings.Fields(text)

	// Step 3: Process each raw token into sub-tokens
	var tokens []Token
	for _, raw := range rawTokens {
		tokens = append(tokens, processRawToken(raw)...)
	}

	// Step 4: Mark sentence-first tokens
	markSentenceFirst(tokens)

	// Step 5: Mark stopwords
	stopwords := data.LoadStopwords()
	for i := range tokens {
		if !tokens[i].IsFiltered {
			tokens[i].IsStopword = stopwords[tokens[i].Lower]
		}
	}

	// Step 6: Detect proper nouns (must run after stopword marking)
	detectProperNouns(tokens)

	return tokens
}

// normalizeAndFilterASCII normalizes common Unicode characters to ASCII
// equivalents and removes remaining non-ASCII characters.
func normalizeAndFilterASCII(text string) string {
	replacer := strings.NewReplacer(
		"\u2019", "'", // ' right single quotation mark
		"\u2018", "'", // ' left single quotation mark
		"\u201C", "\"", // " left double quotation mark
		"\u201D", "\"", // " right double quotation mark
		"\u2014", "-", // — em dash
		"\u2013", "-", // – en dash
		"\u2026", "...", // … horizontal ellipsis
	)
	text = replacer.Replace(text)

	var buf strings.Builder
	buf.Grow(len(text))
	for _, r := range text {
		if r < 128 {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

// processRawToken processes a single whitespace-delimited token into one or
// more Tokens, handling contractions, possessives, hyphens, and punctuation.
func processRawToken(raw string) []Token {
	segments := splitSegments(raw)

	var tokens []Token
	for _, seg := range segments {
		if len(seg) == 0 {
			continue
		}

		// Non-word segment → punctuation token
		if !isWordByte(seg[0]) {
			tokens = append(tokens, Token{
				Original:   seg,
				Lower:      seg,
				IsFiltered: true,
			})
			continue
		}

		// Word segment: handle contractions, then split hyphens
		word := handleContraction(seg)

		parts := splitHyphens(word)
		for _, p := range parts {
			p = strings.Trim(p, "'")
			if p == "" {
				continue
			}

			t := Token{
				Original: p,
				Lower:    strings.ToLower(p),
			}
			if !hasLetter(p) {
				t.IsFiltered = true
			}
			tokens = append(tokens, t)
		}
	}
	return tokens
}

// splitSegments splits a raw token into alternating word-character and
// non-word-character segments. Word characters are ASCII letters, digits,
// apostrophes, and hyphens.
func splitSegments(s string) []string {
	if s == "" {
		return nil
	}

	var segments []string
	i := 0
	for i < len(s) {
		if isWordByte(s[i]) {
			j := i + 1
			for j < len(s) && isWordByte(s[j]) {
				j++
			}
			segments = append(segments, s[i:j])
			i = j
		} else {
			j := i + 1
			for j < len(s) && !isWordByte(s[j]) {
				j++
			}
			segments = append(segments, s[i:j])
			i = j
		}
	}
	return segments
}

// isWordByte reports whether b is an ASCII letter, digit, apostrophe, or hyphen.
func isWordByte(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') ||
		(b >= '0' && b <= '9') || b == '\'' || b == '-'
}

// handleContraction processes English contractions and possessives,
// returning the base word with original casing preserved where possible.
func handleContraction(word string) string {
	lower := strings.ToLower(word)

	// Special irregular contractions
	switch lower {
	case "won't":
		return "will"
	case "can't":
		return "can"
	case "shan't":
		return "shall"
	}

	// n't contractions: don't → do, doesn't → does, etc.
	if strings.HasSuffix(lower, "n't") && len(word) > 3 {
		return word[:len(word)-3]
	}

	// 're contractions: they're → they
	if strings.HasSuffix(lower, "'re") && len(word) > 3 {
		return word[:len(word)-3]
	}

	// 've contractions: we've → we
	if strings.HasSuffix(lower, "'ve") && len(word) > 3 {
		return word[:len(word)-3]
	}

	// 'll contractions: she'll → she
	if strings.HasSuffix(lower, "'ll") && len(word) > 3 {
		return word[:len(word)-3]
	}

	// 'm contractions: I'm → I
	if strings.HasSuffix(lower, "'m") && len(word) > 2 {
		return word[:len(word)-2]
	}

	// 'd contractions: I'd → I
	if strings.HasSuffix(lower, "'d") && len(word) > 2 {
		return word[:len(word)-2]
	}

	// 's contractions and possessives: he's → he, author's → author
	if strings.HasSuffix(lower, "'s") && len(word) > 2 {
		return word[:len(word)-2]
	}

	return word
}

// splitHyphens splits a word on hyphens, returning non-empty parts.
func splitHyphens(word string) []string {
	if !strings.Contains(word, "-") {
		return []string{word}
	}

	raw := strings.Split(word, "-")
	parts := make([]string, 0, len(raw))
	for _, p := range raw {
		if p != "" {
			parts = append(parts, p)
		}
	}
	return parts
}

// markSentenceFirst marks the first non-filtered token and each
// non-filtered token following a sentence-ending punctuation (.!?) as IsFirst.
func markSentenceFirst(tokens []Token) {
	expectFirst := true
	for i := range tokens {
		if tokens[i].IsFiltered {
			if isSentenceEnding(tokens[i].Original) {
				expectFirst = true
			}
			continue
		}
		if expectFirst {
			tokens[i].IsFirst = true
			expectFirst = false
		}
	}
}

// isSentenceEnding reports whether the token contains sentence-ending punctuation.
func isSentenceEnding(s string) bool {
	return strings.ContainsAny(s, ".!?")
}

// isStopword reports whether the lowercase word is in the stopword list.
func isStopword(word string) bool {
	return data.LoadStopwords()[word]
}

// detectProperNouns marks tokens as proper nouns based on capitalization patterns.
// Prerequisites: IsStopword must be set before calling this function.
func detectProperNouns(tokens []Token) {
	// Pass 1: Mark obvious proper nouns
	for i := range tokens {
		if tokens[i].IsFiltered || tokens[i].IsStopword {
			continue
		}

		orig := tokens[i].Original
		if len(orig) == 0 {
			continue
		}

		// All uppercase with length > 1 → acronym (e.g., NASA, FBI)
		if len(orig) > 1 && isAllUpper(orig) {
			tokens[i].IsProper = true
			continue
		}

		// Non-sentence-first word starting with uppercase → proper noun
		if !tokens[i].IsFirst && isUpperFirst(orig) {
			tokens[i].IsProper = true
		}
	}

	// Pass 2: Handle consecutive capitalized words (FR-207)
	// Non-sentence-first capitalized sequences form proper noun phrases.
	// Sentence-first tokens break the run to avoid false positives
	// (e.g., "Visit New York" — "Visit" is a verb, not part of the proper noun).
	i := 0
	for i < len(tokens) {
		// Sentence-first, filtered, stopword, or non-capitalized tokens break the run
		if tokens[i].IsFiltered || tokens[i].IsStopword || tokens[i].IsFirst || !isUpperFirst(tokens[i].Original) {
			i++
			continue
		}

		// Start of a non-sentence-first capitalized run
		start := i
		for i < len(tokens) && !tokens[i].IsFiltered && !tokens[i].IsStopword &&
			!tokens[i].IsFirst && len(tokens[i].Original) > 0 && isUpperFirst(tokens[i].Original) {
			i++
		}

		// Mark all tokens in the run as proper nouns
		if (i - start) > 1 {
			for j := start; j < i; j++ {
				tokens[j].IsProper = true
			}
		}
	}
}

// isAllUpper reports whether all characters in s are uppercase letters.
func isAllUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

// isUpperFirst reports whether the first character of s is an uppercase letter.
func isUpperFirst(s string) bool {
	if s == "" {
		return false
	}
	return unicode.IsUpper(rune(s[0]))
}

// hasLetter reports whether s contains at least one ASCII letter.
func hasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// splitSentences splits English text into individual sentences.
// It handles abbreviations, ellipses, and simplified quoted-period detection.
// Returns nil for empty or whitespace-only input.
func splitSentences(text string) []string {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil
	}

	abbrevs := data.LoadAbbreviations()

	var sentences []string
	var current strings.Builder
	inQuote := false

	i := 0
	for i < len(text) {
		ch := text[i]

		// Track quote state
		if ch == '"' {
			inQuote = !inQuote
			current.WriteByte(ch)
			i++
			continue
		}

		// Check for ellipsis (three consecutive dots)
		if ch == '.' && i+2 < len(text) && text[i+1] == '.' && text[i+2] == '.' {
			current.WriteString("...")
			i += 3
			// Ellipsis acts as a sentence break
			s := strings.TrimSpace(current.String())
			if s != "" {
				sentences = append(sentences, s)
			}
			current.Reset()
			continue
		}

		// Sentence-ending punctuation
		if ch == '.' || ch == '!' || ch == '?' {
			current.WriteByte(ch)

			if ch == '.' {
				// Check if dot is part of a known abbreviation
				if sentenceIsAbbreviation(text, i, abbrevs) {
					if !abbreviationEndsSentence(text, i) {
						i++
						continue
					}
				}
				// Inside quotes: only split if followed by (optional quote +) space + uppercase
				if inQuote && !followedBySpaceUppercase(text, i) {
					i++
					continue
				}
			}

			i++

			// Consume additional consecutive sentence-ending punctuation (e.g., ?!)
			for i < len(text) && (text[i] == '!' || text[i] == '?' || text[i] == '.') {
				current.WriteByte(text[i])
				i++
			}

			// Consume closing quote(s) immediately after punctuation
			for i < len(text) && text[i] == '"' {
				current.WriteByte(text[i])
				inQuote = !inQuote
				i++
			}

			s := strings.TrimSpace(current.String())
			if s != "" {
				sentences = append(sentences, s)
			}
			current.Reset()
			continue
		}

		current.WriteByte(ch)
		i++
	}

	// Remaining text after last split point
	s := strings.TrimSpace(current.String())
	if s != "" {
		sentences = append(sentences, s)
	}

	if len(sentences) == 0 {
		return nil
	}
	return sentences
}

// sentenceIsAbbreviation checks whether the dot at position dotPos in text
// terminates a known abbreviation. It scans backward and forward for letters
// and dots to handle multi-dot abbreviations like "i.e." and "e.g.".
func sentenceIsAbbreviation(text string, dotPos int, abbrevs map[string]bool) bool {
	// Scan backward for letters and dots
	start := dotPos
	for start > 0 {
		c := text[start-1]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '.' {
			start--
		} else {
			break
		}
	}
	if start == dotPos {
		return false // no letters before the dot
	}

	// Scan forward for letters and dots (multi-dot abbreviations like i.e.)
	end := dotPos + 1
	for end < len(text) {
		c := text[end]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '.' {
			end++
		} else {
			break
		}
	}

	// Check full span first (e.g., "i.e." when at either dot)
	fullWord := strings.ToLower(text[start:end])
	if abbrevs[fullWord] {
		return true
	}

	// Check up to current dot only (e.g., "mr.")
	word := strings.ToLower(text[start : dotPos+1])
	return abbrevs[word]
}

// followedBySpaceUppercase reports whether position pos in text is followed
// by an optional closing quote, then a space, then an uppercase letter.
func followedBySpaceUppercase(text string, pos int) bool {
	j := pos + 1
	// Skip closing quote(s)
	for j < len(text) && text[j] == '"' {
		j++
	}
	// Expect space then uppercase letter
	if j < len(text) && text[j] == ' ' {
		j++
		if j < len(text) && text[j] >= 'A' && text[j] <= 'Z' {
			return true
		}
	}
	return false
}

// abbreviationEndsSentence reports whether an abbreviation-ending dot at dotPos
// should be treated as a sentence boundary.
func abbreviationEndsSentence(text string, dotPos int) bool {
	abbrev := abbreviationAtDot(text, dotPos)
	j := dotPos + 1

	// Optional closing quote(s) immediately after dot.
	for j < len(text) && text[j] == '"' {
		j++
	}

	if j >= len(text) {
		return true
	}

	if text[j] == '\n' || text[j] == '\r' {
		return true
	}

	// If next char is not whitespace, abbreviation is mid-sentence.
	if !unicode.IsSpace(rune(text[j])) {
		return false
	}

	// Consume whitespace; any newline in this run is sentence end.
	for j < len(text) && unicode.IsSpace(rune(text[j])) {
		if text[j] == '\n' || text[j] == '\r' {
			return true
		}
		j++
	}

	// Optional quote(s) before next token.
	for j < len(text) && text[j] == '"' {
		j++
	}

	if j >= len(text) {
		return true
	}

	r, _ := utf8.DecodeRuneInString(text[j:])
	if !unicode.IsUpper(r) {
		return false
	}
	return !isTitleAbbreviation(abbrev)
}

// abbreviationAtDot returns the lowercase abbreviation token ending at dotPos,
// such as "mr." or "i.e.".
func abbreviationAtDot(text string, dotPos int) string {
	start := dotPos
	for start > 0 {
		c := text[start-1]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '.' {
			start--
			continue
		}
		break
	}
	if start > dotPos {
		return ""
	}
	return strings.ToLower(text[start : dotPos+1])
}

// isTitleAbbreviation reports abbreviations that are commonly sentence-internal
// before proper names (e.g., "Mr. Smith").
func isTitleAbbreviation(abbrev string) bool {
	switch abbrev {
	case "mr.", "mrs.", "ms.", "dr.", "prof.", "sr.", "jr.", "st.":
		return true
	default:
		return false
	}
}
