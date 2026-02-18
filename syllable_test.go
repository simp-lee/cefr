package cefr

import (
	"fmt"
	"testing"
)

func TestCountSyllables(t *testing.T) {
	tests := []struct {
		word string
		want int
	}{
		// Minimum-1 rule and edge cases.
		{"", 0},
		{"a", 1},
		{"I", 1},
		{"the", 1},

		// Basic vowel group counting.
		{"cat", 1},
		{"dog", 1},
		{"run", 1},
		{"go", 1},
		{"meet", 1},
		{"book", 1},
		{"school", 1},

		// Multi-syllable words.
		{"beautiful", 3},
		{"animal", 3},
		{"elephant", 3},
		{"important", 3},
		{"education", 4},

		// Silent -e (not -le).
		{"smile", 1},
		{"make", 1},
		{"fine", 1},
		{"home", 1},
		{"like", 1},
		{"name", 1},

		// Consonant + le (syllabic -le).
		{"table", 2},
		{"simple", 2},
		{"little", 2},
		{"people", 2},
		{"purple", 2},
		{"castle", 2},
		{"cycle", 2},
		{"bottle", 2},

		// Vowel + le → silent e.
		{"pale", 1},
		{"whale", 1},
		{"scale", 1},

		// Trailing -ed: preceded by t/d → syllable.
		{"wanted", 2},
		{"added", 2},
		{"started", 2},
		{"needed", 2},

		// Trailing -ed: preceded by other → silent.
		{"walked", 1},
		{"played", 1},
		{"liked", 1},
		{"jumped", 1},
		{"asked", 1},

		// Trailing -es: preceded by sibilant (s, x, z, ch, sh) → syllable.
		{"boxes", 2},
		{"watches", 2},
		{"dishes", 2},
		{"fixes", 2},
		{"buzzes", 2},
		{"cases", 2},

		// Trailing -es: preceded by other → silent.
		{"makes", 1},
		{"likes", 1},
		{"games", 1},
		{"comes", 1},

		// Words with y as vowel.
		{"my", 1},
		{"fly", 1},
		{"happy", 2},
		{"mystery", 3},

		// Double vowels counted as one group.
		{"three", 1},
		{"free", 1},
		{"eye", 1},

		// Additional trailing -e words.
		{"there", 1},

		// Trailing -es: non-sibilant → silent.
		{"goes", 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s=%d", tt.word, tt.want), func(t *testing.T) {
			got := countSyllables(tt.word)
			if got != tt.want {
				t.Errorf("countSyllables(%q) = %d; want %d", tt.word, got, tt.want)
			}
		})
	}
}

func TestCountTotalSyllables(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{"nil", nil, 0},
		{"empty", []string{}, 0},
		{"single", []string{"cat"}, 1},
		{"two words", []string{"cat", "dog"}, 2},
		{"mixed", []string{"beautiful", "cat", "table"}, 6}, // 3+1+2
		{"sentence words", []string{"the", "cat", "sat", "on", "the", "mat"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countTotalSyllables(tt.words)
			if got != tt.want {
				t.Errorf("countTotalSyllables(%v) = %d; want %d", tt.words, got, tt.want)
			}
		})
	}
}
