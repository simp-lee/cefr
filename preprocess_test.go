package cefr

import (
	"testing"
)

func TestHandleContraction(t *testing.T) {
	tests := []struct {
		name string
		word string
		want string
	}{
		// n't contractions
		{name: "don't", word: "don't", want: "do"},
		{name: "doesn't", word: "doesn't", want: "does"},
		{name: "didn't", word: "didn't", want: "did"},
		{name: "isn't", word: "isn't", want: "is"},
		{name: "aren't", word: "aren't", want: "are"},
		{name: "wasn't", word: "wasn't", want: "was"},
		{name: "weren't", word: "weren't", want: "were"},
		{name: "hasn't", word: "hasn't", want: "has"},
		{name: "haven't", word: "haven't", want: "have"},
		{name: "couldn't", word: "couldn't", want: "could"},
		{name: "wouldn't", word: "wouldn't", want: "would"},
		{name: "shouldn't", word: "shouldn't", want: "should"},

		// Special irregular contractions
		{name: "won't", word: "won't", want: "will"},
		{name: "can't", word: "can't", want: "can"},
		{name: "shan't", word: "shan't", want: "shall"},

		// Case-preserved n't
		{name: "Don't", word: "Don't", want: "Do"},
		{name: "Won't uppercase", word: "Won't", want: "will"},
		{name: "Can't uppercase", word: "Can't", want: "can"},

		// 'm contraction
		{name: "I'm", word: "I'm", want: "I"},

		// 're contractions
		{name: "they're", word: "they're", want: "they"},
		{name: "we're", word: "we're", want: "we"},
		{name: "you're", word: "you're", want: "you"},

		// 've contractions
		{name: "we've", word: "we've", want: "we"},
		{name: "I've", word: "I've", want: "I"},
		{name: "they've", word: "they've", want: "they"},

		// 'd contractions
		{name: "I'd", word: "I'd", want: "I"},
		{name: "he'd", word: "he'd", want: "he"},

		// 'll contractions
		{name: "she'll", word: "she'll", want: "she"},
		{name: "I'll", word: "I'll", want: "I"},
		{name: "they'll", word: "they'll", want: "they"},

		// 's contractions/possessives
		{name: "he's", word: "he's", want: "he"},
		{name: "she's", word: "she's", want: "she"},
		{name: "it's", word: "it's", want: "it"},
		{name: "author's", word: "author's", want: "author"},
		{name: "that's", word: "that's", want: "that"},

		// No contraction
		{name: "plain word", word: "hello", want: "hello"},
		{name: "no apostrophe", word: "world", want: "world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := handleContraction(tt.word)
			if got != tt.want {
				t.Errorf("handleContraction(%q) = %q, want %q", tt.word, got, tt.want)
			}
		})
	}
}

func TestNormalizeAndFilterASCII(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "plain ASCII", input: "Hello world", want: "Hello world"},
		{name: "smart single quotes", input: "can\u2019t", want: "can't"},
		{name: "smart double quotes", input: "\u201CHello\u201D", want: "\"Hello\""},
		{name: "em dash", input: "word\u2014word", want: "word-word"},
		{name: "en dash", input: "1\u20132", want: "1-2"},
		{name: "ellipsis", input: "wait\u2026", want: "wait..."},
		{name: "Chinese characters removed", input: "Hello 你好 world", want: "Hello  world"},
		{name: "mixed non-ASCII", input: "café résumé", want: "caf rsum"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeAndFilterASCII(tt.input)
			if got != tt.want {
				t.Errorf("normalizeAndFilterASCII(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestSplitSegments(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{name: "plain word", input: "hello", want: []string{"hello"}},
		{name: "word with comma", input: "hello,", want: []string{"hello", ","}},
		{name: "parenthesized", input: "(hello)", want: []string{"(", "hello", ")"}},
		{name: "contraction", input: "can't", want: []string{"can't"}},
		{name: "hyphenated", input: "well-known", want: []string{"well-known"}},
		{name: "ellipsis after", input: "word...", want: []string{"word", "..."}},
		{name: "empty string", input: "", want: nil},
		{name: "pure punctuation", input: "...", want: []string{"..."}},
		{name: "complex", input: "(well-known)", want: []string{"(", "well-known", ")"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitSegments(tt.input)
			if !stringSliceEqual(got, tt.want) {
				t.Errorf("splitSegments(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestProcessRawToken(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		wantOriginals []string
		wantFiltered  []bool
	}{
		{
			name:          "plain word",
			input:         "hello",
			wantOriginals: []string{"hello"},
			wantFiltered:  []bool{false},
		},
		{
			name:          "word with trailing comma",
			input:         "hello,",
			wantOriginals: []string{"hello", ","},
			wantFiltered:  []bool{false, true},
		},
		{
			name:          "contraction",
			input:         "can't",
			wantOriginals: []string{"can"},
			wantFiltered:  []bool{false},
		},
		{
			name:          "won't special",
			input:         "won't",
			wantOriginals: []string{"will"},
			wantFiltered:  []bool{false},
		},
		{
			name:          "hyphenated",
			input:         "well-known",
			wantOriginals: []string{"well", "known"},
			wantFiltered:  []bool{false, false},
		},
		{
			name:          "possessive with comma",
			input:         "author's,",
			wantOriginals: []string{"author", ","},
			wantFiltered:  []bool{false, true},
		},
		{
			name:          "number",
			input:         "100",
			wantOriginals: []string{"100"},
			wantFiltered:  []bool{true},
		},
		{
			name:          "period",
			input:         ".",
			wantOriginals: []string{"."},
			wantFiltered:  []bool{true},
		},
		{
			name:          "parenthesized word",
			input:         "(hello)",
			wantOriginals: []string{"(", "hello", ")"},
			wantFiltered:  []bool{true, false, true},
		},
		{
			name:          "word with period",
			input:         "world.",
			wantOriginals: []string{"world", "."},
			wantFiltered:  []bool{false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := processRawToken(tt.input)
			if len(got) != len(tt.wantOriginals) {
				t.Fatalf("processRawToken(%q) returned %d tokens, want %d: %v",
					tt.input, len(got), len(tt.wantOriginals), tokensOriginals(got))
			}
			for i := range got {
				if got[i].Original != tt.wantOriginals[i] {
					t.Errorf("token[%d].Original = %q, want %q", i, got[i].Original, tt.wantOriginals[i])
				}
				if got[i].IsFiltered != tt.wantFiltered[i] {
					t.Errorf("token[%d].IsFiltered = %v, want %v (Original=%q)",
						i, got[i].IsFiltered, tt.wantFiltered[i], got[i].Original)
				}
			}
		})
	}
}

func TestMarkSentenceFirst(t *testing.T) {
	tests := []struct {
		name   string
		tokens []Token
		want   []bool // expected IsFirst values
	}{
		{
			name: "first token is sentence-first",
			tokens: []Token{
				{Original: "Hello", Lower: "hello"},
				{Original: "world", Lower: "world"},
			},
			want: []bool{true, false},
		},
		{
			name: "after period",
			tokens: []Token{
				{Original: "Hello", Lower: "hello"},
				{Original: ".", Lower: ".", IsFiltered: true},
				{Original: "World", Lower: "world"},
			},
			want: []bool{true, false, true},
		},
		{
			name: "skip filtered for sentence-first",
			tokens: []Token{
				{Original: ".", Lower: ".", IsFiltered: true},
				{Original: "100", Lower: "100", IsFiltered: true},
				{Original: "Hello", Lower: "hello"},
			},
			want: []bool{false, false, true},
		},
		{
			name: "multiple sentences",
			tokens: []Token{
				{Original: "Hi", Lower: "hi"},
				{Original: ".", Lower: ".", IsFiltered: true},
				{Original: "Bye", Lower: "bye"},
				{Original: "!", Lower: "!", IsFiltered: true},
				{Original: "What", Lower: "what"},
				{Original: "?", Lower: "?", IsFiltered: true},
				{Original: "Done", Lower: "done"},
			},
			want: []bool{true, false, true, false, true, false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			markSentenceFirst(tt.tokens)
			for i, tok := range tt.tokens {
				if tok.IsFirst != tt.want[i] {
					t.Errorf("token[%d] (%q) IsFirst = %v, want %v",
						i, tok.Original, tok.IsFirst, tt.want[i])
				}
			}
		})
	}
}

func TestIsStopword(t *testing.T) {
	tests := []struct {
		word string
		want bool
	}{
		{"i", true},
		{"the", true},
		{"is", true},
		{"on", true},
		{"hello", true},
		{"computer", false},
		{"beautiful", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			if got := isStopword(tt.word); got != tt.want {
				t.Errorf("isStopword(%q) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}

func TestDetectProperNouns(t *testing.T) {
	tests := []struct {
		name   string
		tokens []Token
		want   []bool // expected IsProper values
	}{
		{
			name: "non-sentence-first capitalized",
			tokens: []Token{
				{Original: "I", Lower: "i", IsFirst: true, IsStopword: true},
				{Original: "went", Lower: "went"},
				{Original: "to", Lower: "to", IsStopword: true},
				{Original: "Paris", Lower: "paris"},
			},
			want: []bool{false, false, false, true},
		},
		{
			name: "all uppercase acronym",
			tokens: []Token{
				{Original: "The", Lower: "the", IsFirst: true, IsStopword: true},
				{Original: "NASA", Lower: "nasa"},
				{Original: "program", Lower: "program"},
			},
			want: []bool{false, true, false},
		},
		{
			name: "I not marked as proper (stopword)",
			tokens: []Token{
				{Original: "I", Lower: "i", IsFirst: true, IsStopword: true},
				{Original: "am", Lower: "am", IsStopword: true},
				{Original: "here", Lower: "here", IsStopword: true},
			},
			want: []bool{false, false, false},
		},
		{
			name: "sentence-first not proper alone",
			tokens: []Token{
				{Original: "The", Lower: "the", IsFirst: true, IsStopword: true},
				{Original: "cat", Lower: "cat"},
				{Original: "is", Lower: "is", IsStopword: true},
				{Original: "cute", Lower: "cute"},
			},
			want: []bool{false, false, false, false},
		},
		{
			name: "consecutive capitalized - FR-207 mid-sentence",
			tokens: []Token{
				{Original: "Visit", Lower: "visit", IsFirst: true},
				{Original: "New", Lower: "new"},
				{Original: "York", Lower: "york"},
				{Original: "today", Lower: "today"},
			},
			// "Visit" is sentence-first → not marked in pass 1, breaks pass 2 run.
			// "New" and "York" are non-sentence-first + capitalized → proper in pass 1.
			want: []bool{false, true, true, false},
		},
		{
			name: "consecutive capitalized mid-sentence",
			tokens: []Token{
				{Original: "I", Lower: "i", IsFirst: true, IsStopword: true},
				{Original: "visited", Lower: "visited"},
				{Original: "New", Lower: "new"},
				{Original: "York", Lower: "york"},
				{Original: "City", Lower: "city"},
			},
			// All three are non-sentence-first + capitalized → all proper
			want: []bool{false, false, true, true, true},
		},
		{
			name: "stopword breaks capitalized run",
			tokens: []Token{
				{Original: "I", Lower: "i", IsFirst: true, IsStopword: true},
				{Original: "like", Lower: "like"},
				{Original: "United", Lower: "united"},
				{Original: "States", Lower: "states"},
				{Original: "of", Lower: "of", IsStopword: true},
				{Original: "America", Lower: "america"},
			},
			// "United" and "States" are non-sentence-first + capitalized → proper
			// "of" is stopword, breaks the run
			// "America" is non-sentence-first + capitalized → proper
			want: []bool{false, false, true, true, false, true},
		},
		{
			name: "single letter not all-uppercase",
			tokens: []Token{
				{Original: "A", Lower: "a", IsFirst: true, IsStopword: true},
				{Original: "cat", Lower: "cat"},
			},
			// "A" is a stopword, so not proper. Also len == 1 so not all-uppercase acronym.
			want: []bool{false, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			detectProperNouns(tt.tokens)
			for i, tok := range tt.tokens {
				if tok.IsProper != tt.want[i] {
					t.Errorf("token[%d] (%q) IsProper = %v, want %v",
						i, tok.Original, tok.IsProper, tt.want[i])
				}
			}
		})
	}
}

func TestTokenize_Integration(t *testing.T) {
	tests := []struct {
		name  string
		input string
		check func(t *testing.T, tokens []Token)
	}{
		{
			name:  "basic sentence",
			input: "The cat sat on the mat.",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				// Expect: The, cat, sat, on, the, mat, .
				originals := tokensOriginals(tokens)
				wantOrig := []string{"The", "cat", "sat", "on", "the", "mat", "."}
				if !stringSliceEqual(originals, wantOrig) {
					t.Errorf("originals = %v, want %v", originals, wantOrig)
				}
				// "The" should be sentence-first
				if !tokens[0].IsFirst {
					t.Error("tokens[0] (The) should be IsFirst")
				}
				// "the", "on" should be stopwords
				if !tokens[0].IsStopword {
					t.Error("tokens[0] (The) should be IsStopword")
				}
				if !tokens[3].IsStopword {
					t.Error("tokens[3] (on) should be IsStopword")
				}
			},
		},
		{
			name:  "contractions and proper nouns",
			input: "I can't go to New York.",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				originals := tokensOriginals(tokens)
				wantOrig := []string{"I", "can", "go", "to", "New", "York", "."}
				if !stringSliceEqual(originals, wantOrig) {
					t.Errorf("originals = %v, want %v", originals, wantOrig)
				}
				// "I" is stopword, not proper
				if !tokens[0].IsStopword {
					t.Error("tokens[0] (I) should be IsStopword")
				}
				if tokens[0].IsProper {
					t.Error("tokens[0] (I) should NOT be IsProper")
				}
				// "New" and "York" are proper nouns
				if !tokens[4].IsProper {
					t.Error("tokens[4] (New) should be IsProper")
				}
				if !tokens[5].IsProper {
					t.Error("tokens[5] (York) should be IsProper")
				}
			},
		},
		{
			name:  "hyphens and numbers",
			input: "The well-known 100 facts.",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				originals := tokensOriginals(tokens)
				wantOrig := []string{"The", "well", "known", "100", "facts", "."}
				if !stringSliceEqual(originals, wantOrig) {
					t.Errorf("originals = %v, want %v", originals, wantOrig)
				}
				// "100" should be filtered
				if !tokens[3].IsFiltered {
					t.Error("tokens[3] (100) should be IsFiltered")
				}
			},
		},
		{
			name:  "multiple sentences",
			input: "Hello world. Good morning! How are you?",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				// Find sentence-first tokens
				var firsts []string
				for _, tok := range tokens {
					if tok.IsFirst {
						firsts = append(firsts, tok.Original)
					}
				}
				wantFirsts := []string{"Hello", "Good", "How"}
				if !stringSliceEqual(firsts, wantFirsts) {
					t.Errorf("sentence-first tokens = %v, want %v", firsts, wantFirsts)
				}
			},
		},
		{
			name:  "non-ASCII filtered",
			input: "Hello 你好 world café",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				originals := tokensOriginals(tokens)
				wantOrig := []string{"Hello", "world", "caf"}
				if !stringSliceEqual(originals, wantOrig) {
					t.Errorf("originals = %v, want %v", originals, wantOrig)
				}
			},
		},
		{
			name:  "smart quotes normalized",
			input: "I can\u2019t believe it\u2019s real.",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				originals := tokensOriginals(tokens)
				wantOrig := []string{"I", "can", "believe", "it", "real", "."}
				if !stringSliceEqual(originals, wantOrig) {
					t.Errorf("originals = %v, want %v", originals, wantOrig)
				}
			},
		},
		{
			name:  "empty input",
			input: "",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				if len(tokens) != 0 {
					t.Errorf("empty input should produce 0 tokens, got %d", len(tokens))
				}
			},
		},
		{
			name:  "only whitespace",
			input: "   \t\n  ",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				if len(tokens) != 0 {
					t.Errorf("whitespace-only input should produce 0 tokens, got %d", len(tokens))
				}
			},
		},
		{
			name:  "lowercase preserved in Lower field",
			input: "Hello WORLD",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				if tokens[0].Lower != "hello" {
					t.Errorf("tokens[0].Lower = %q, want %q", tokens[0].Lower, "hello")
				}
				if tokens[1].Lower != "world" {
					t.Errorf("tokens[1].Lower = %q, want %q", tokens[1].Lower, "world")
				}
			},
		},
		{
			name:  "NASA all-uppercase proper noun",
			input: "I work at NASA today.",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				// Find NASA token
				for _, tok := range tokens {
					if tok.Original == "NASA" {
						if !tok.IsProper {
							t.Error("NASA should be IsProper")
						}
						return
					}
				}
				t.Error("NASA token not found")
			},
		},
		{
			name:  "possessive stripping",
			input: "The author's book is great.",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				originals := tokensOriginals(tokens)
				wantOrig := []string{"The", "author", "book", "is", "great", "."}
				if !stringSliceEqual(originals, wantOrig) {
					t.Errorf("originals = %v, want %v", originals, wantOrig)
				}
			},
		},
		{
			name:  "won't special case",
			input: "She won't go.",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				originals := tokensOriginals(tokens)
				wantOrig := []string{"She", "will", "go", "."}
				if !stringSliceEqual(originals, wantOrig) {
					t.Errorf("originals = %v, want %v", originals, wantOrig)
				}
			},
		},
		{
			name:  "mother-in-law triple hyphen",
			input: "My mother-in-law arrived.",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				originals := tokensOriginals(tokens)
				wantOrig := []string{"My", "mother", "in", "law", "arrived", "."}
				if !stringSliceEqual(originals, wantOrig) {
					t.Errorf("originals = %v, want %v", originals, wantOrig)
				}
			},
		},
		{
			name:  "children possessive",
			input: "The children's toys are here.",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				originals := tokensOriginals(tokens)
				wantOrig := []string{"The", "children", "toys", "are", "here", "."}
				if !stringSliceEqual(originals, wantOrig) {
					t.Errorf("originals = %v, want %v", originals, wantOrig)
				}
			},
		},
		{
			name:  "multiple contractions",
			input: "I'm happy, they're here, we've gone.",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				originals := tokensOriginals(tokens)
				wantOrig := []string{"I", "happy", ",", "they", "here", ",", "we", "gone", "."}
				if !stringSliceEqual(originals, wantOrig) {
					t.Errorf("originals = %v, want %v", originals, wantOrig)
				}
			},
		},
		{
			name:  "contractions I'd and she'll",
			input: "I'd say she'll come.",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				originals := tokensOriginals(tokens)
				wantOrig := []string{"I", "say", "she", "come", "."}
				if !stringSliceEqual(originals, wantOrig) {
					t.Errorf("originals = %v, want %v", originals, wantOrig)
				}
			},
		},
		{
			name:  "pure punctuation tokens filtered",
			input: "Hello !!! world",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				for _, tok := range tokens {
					if tok.Original == "!!!" && !tok.IsFiltered {
						t.Error("pure punctuation '!!!' should be IsFiltered")
					}
				}
			},
		},
		{
			name:  "decimal number tokens all filtered",
			input: "The price is 45.6 dollars.",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				for _, tok := range tokens {
					if tok.Original == "45" || tok.Original == "6" {
						if !tok.IsFiltered {
							t.Errorf("numeric token %q should be IsFiltered", tok.Original)
						}
					}
				}
				// The dot between 45 and 6 should also be filtered
				var foundNumDot bool
				for i, tok := range tokens {
					if tok.Original == "." && i > 0 && tokens[i-1].Original == "45" {
						foundNumDot = true
						if !tok.IsFiltered {
							t.Error("dot in decimal number should be IsFiltered")
						}
					}
				}
				if !foundNumDot {
					t.Error("expected dot token between 45 and 6")
				}
			},
		},
		{
			name:  "all uppercase stopwords not proper",
			input: "WE ARE ON IT",
			check: func(t *testing.T, tokens []Token) {
				t.Helper()
				for _, tok := range tokens {
					if tok.IsProper {
						t.Errorf("token %q should NOT be IsProper (all-uppercase stopword text)", tok.Original)
					}
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := tokenize(tt.input)
			tt.check(t, tokens)
		})
	}
}

// --- helpers ---

func tokensOriginals(tokens []Token) []string {
	out := make([]string, len(tokens))
	for i, t := range tokens {
		out[i] = t.Original
	}
	return out
}

func TestSplitSentences(t *testing.T) {
	tests := []struct {
		name string
		text string
		want []string
	}{
		// Empty / whitespace
		{name: "empty string", text: "", want: nil},
		{name: "whitespace only", text: "   ", want: nil},

		// Basic splitting
		{name: "single sentence with period", text: "Hello world.", want: []string{"Hello world."}},
		{name: "single sentence no ending punct", text: "Hello world", want: []string{"Hello world"}},
		{name: "two sentences period", text: "Hello. World.", want: []string{"Hello.", "World."}},
		{name: "exclamation mark", text: "Wow! That is great.", want: []string{"Wow!", "That is great."}},
		{name: "question mark", text: "How are you? I am fine.", want: []string{"How are you?", "I am fine."}},
		{name: "mixed punctuation", text: "Hello. Really? Yes!", want: []string{"Hello.", "Really?", "Yes!"}},

		// Abbreviations — should NOT split
		{name: "abbreviation Mr", text: "Mr. Smith went home.", want: []string{"Mr. Smith went home."}},
		{name: "abbreviation Dr", text: "Dr. Jones is here.", want: []string{"Dr. Jones is here."}},
		{name: "abbreviation Mrs", text: "Mrs. Brown left.", want: []string{"Mrs. Brown left."}},
		{name: "abbreviation Prof", text: "Prof. Lee teaches math.", want: []string{"Prof. Lee teaches math."}},
		{name: "abbreviation etc", text: "We have cats, dogs, etc. at the farm.", want: []string{"We have cats, dogs, etc. at the farm."}},
		{name: "abbreviation etc at sentence end should split", text: "We used tools, etc. This is enough.", want: []string{"We used tools, etc.", "This is enough."}},
		{name: "abbreviation e.g.", text: "Use tools e.g. a hammer.", want: []string{"Use tools e.g. a hammer."}},
		{name: "abbreviation i.e.", text: "That is i.e. correct.", want: []string{"That is i.e. correct."}},
		{name: "abbreviation case insensitive", text: "MR. Smith went home.", want: []string{"MR. Smith went home."}},

		// Ellipsis
		{name: "ellipsis mid text", text: "Wait... What happened?", want: []string{"Wait...", "What happened?"}},
		{name: "ellipsis at end", text: "I wonder...", want: []string{"I wonder..."}},
		{name: "ellipsis only", text: "...", want: []string{"..."}},

		// Quoted periods
		{name: "quoted period no split", text: `He said "I love N.Y. very much."`, want: []string{`He said "I love N.Y. very much."`}},
		{name: "quoted sentence split after close", text: `"Hello." She left.`, want: []string{`"Hello."`, `She left.`}},
		{name: "quoted with continuation", text: `He said "That was good." Then he left.`, want: []string{`He said "That was good."`, `Then he left.`}},

		// Consecutive punctuation
		{name: "question exclamation", text: "Really?! Yes.", want: []string{"Really?!", "Yes."}},

		// Whitespace trimming
		{name: "leading trailing whitespace", text: "  Hello.  World.  ", want: []string{"Hello.", "World."}},

		// Complex / mixed
		{name: "mixed abbreviations and sentences", text: "Mr. Smith said hello. Dr. Jones replied! What happened?",
			want: []string{"Mr. Smith said hello.", "Dr. Jones replied!", "What happened?"}},
		{name: "sentence without ending punctuation at end", text: "First sentence. Second part",
			want: []string{"First sentence.", "Second part"}},

		// Ellipsis followed by lowercase continuation
		{name: "ellipsis then lowercase", text: "He said... and left.",
			want: []string{"He said...", "and left."}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitSentences(tt.text)
			if !stringSliceEqual(got, tt.want) {
				t.Errorf("splitSentences(%q)\n got  %q\n want %q", tt.text, got, tt.want)
			}
		})
	}
}

func stringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
