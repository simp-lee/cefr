package cefr

import "testing"

func TestLemmatize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Irregular forms
		{name: "irregular verb past", input: "went", want: "go"},
		{name: "irregular adjective comp", input: "better", want: "good"},
		{name: "irregular noun plural", input: "children", want: "child"},
		{name: "irregular noun plural mice", input: "mice", want: "mouse"},

		// -ing: doubled consonant
		{name: "ing doubled consonant", input: "running", want: "run"},
		{name: "ing doubled consonant stop", input: "stopping", want: "stop"},
		// -ing: restore e
		{name: "ing restore e", input: "making", want: "make"},
		{name: "ing restore e hoping", input: "hoping", want: "hope"},
		// -ing: direct strip
		{name: "ing direct strip", input: "walking", want: "walk"},

		// -ed: ied → y
		{name: "ed ied to y", input: "studied", want: "study"},
		// -ed: doubled consonant
		{name: "ed doubled consonant", input: "stopped", want: "stop"},
		// -ed: restore e
		{name: "ed restore e", input: "hoped", want: "hope"},
		// -ed: direct strip
		{name: "ed direct strip", input: "walked", want: "walk"},

		// -es: ies → y
		{name: "es ies to y", input: "stories", want: "story"},
		// -es: xes → strip es
		{name: "es xes strip", input: "boxes", want: "box"},
		// -es: shes → strip es
		{name: "es shes strip", input: "dishes", want: "dish"},
		// -es: ches → strip es
		{name: "es ches strip", input: "watches", want: "watch"},
		// -es: sses → strip es
		// -s: direct strip
		{name: "s direct strip", input: "makes", want: "make"},
		{name: "s direct strip walks", input: "walks", want: "walk"},
		{name: "s direct strip stops", input: "stops", want: "stop"},
		{name: "s direct strip cats", input: "cats", want: "cat"},

		// -er: ier → y
		{name: "er ier to y", input: "happier", want: "happy"},
		// -er: doubled consonant
		{name: "er doubled consonant", input: "bigger", want: "big"},
		// -er: restore e
		{name: "er restore e nicer", input: "nicer", want: "nice"},
		{name: "er restore e wider", input: "wider", want: "wide"},
		{name: "er restore e larger", input: "larger", want: "large"},
		// -er: direct strip
		// (walker→walk could work, but "walker" might itself be a word)

		// -est: iest → y
		{name: "est iest to y", input: "happiest", want: "happy"},
		// -est: doubled consonant
		{name: "est doubled consonant", input: "biggest", want: "big"},
		// -est: restore e
		{name: "est restore e nicest", input: "nicest", want: "nice"},
		{name: "est restore e widest", input: "widest", want: "wide"},
		{name: "est restore e largest", input: "largest", want: "large"},

		// -ly: ily → y
		{name: "ly ily to y", input: "happily", want: "happy"},
		{name: "ly ily to y angrily", input: "angrily", want: "angry"},
		// -ly: direct strip (gracefully is NOT in vocab, graceful IS)
		{name: "ly direct strip", input: "gracefully", want: "graceful"},

		// Word already in vocab returns itself (quickly is in Oxford 5000)
		{name: "already in vocab quickly", input: "quickly", want: "quickly"},

		// Already a base form
		{name: "base form", input: "run", want: "run"},
		{name: "base form walk", input: "walk", want: "walk"},

		// Unknown word (not in vocab, no valid stripping)
		{name: "unknown word", input: "xyzabc", want: "xyzabc"},

		// Short words should not be over-stripped
		{name: "short word as", input: "as", want: "as"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lemmatize(tt.input)
			if got != tt.want {
				t.Errorf("lemmatize(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsDoubledConsonant(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"runn", true},
		{"stopp", true},
		{"bigg", true},
		{"walk", false},
		{"make", false},
		{"aa", false},  // vowel, not consonant
		{"nn", true},   // two-char doubled consonant
		{"n", false},   // too short
		{"", false},    // empty
		{"bell", true}, // doubled l
		{"miss", true}, // doubled s
		{"buzz", true}, // doubled z
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := isDoubledConsonant(tt.input)
			if got != tt.want {
				t.Errorf("isDoubledConsonant(%q) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsInVocab(t *testing.T) {
	tests := []struct {
		word string
		want bool
	}{
		{"run", true},
		{"walk", true},
		{"xyzabc", false},
		{"make", true},
		{"the", true},
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			got := isInVocab(tt.word)
			if got != tt.want {
				t.Errorf("isInVocab(%q) = %v; want %v", tt.word, got, tt.want)
			}
		})
	}
}
