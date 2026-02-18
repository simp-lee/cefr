package data

import (
	"testing"
)

func TestLoadOxford(t *testing.T) {
	m := LoadOxford()
	if len(m) == 0 {
		t.Fatal("LoadOxford() returned empty map")
	}
	if len(m) < 4000 {
		t.Errorf("LoadOxford() has %d entries; want at least 4000", len(m))
	}

	// Spot-check a few well-known words
	tests := []struct {
		word string
		want int
	}{
		{"the", 1},
		{"book", 1},
		{"cat", 1},
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			got, ok := m[tt.word]
			if !ok {
				t.Errorf("word %q not found in Oxford map", tt.word)
				return
			}
			if got != tt.want {
				t.Errorf("Oxford[%q] = %d; want %d", tt.word, got, tt.want)
			}
		})
	}
}

func TestLoadOxfordLevelRange(t *testing.T) {
	m := LoadOxford()
	for word, level := range m {
		if level < 1 || level > 5 {
			t.Errorf("Oxford[%q] = %d; want 1-5", word, level)
		}
	}
}

func TestLoadOxfordAllLevelsPresent(t *testing.T) {
	m := LoadOxford()
	counts := make(map[int]int)
	for _, level := range m {
		counts[level]++
	}
	for level := 1; level <= 5; level++ {
		if counts[level] == 0 {
			t.Errorf("no words at Oxford level %d", level)
		}
	}
	t.Logf("Oxford level distribution: %v", counts)
}

func TestLoadNGSL(t *testing.T) {
	m := LoadNGSL()
	if len(m) == 0 {
		t.Fatal("LoadNGSL() returned empty map")
	}
	if len(m) < 2500 {
		t.Errorf("LoadNGSL() has %d entries; want at least 2500", len(m))
	}
	if got, ok := m["the"]; !ok {
		t.Error("word \"the\" not found in NGSL map")
	} else if got != 1 {
		t.Errorf("NGSL[\"the\"] = %d; want 1", got)
	}
}

func TestLoadNGSLLevelRange(t *testing.T) {
	m := LoadNGSL()
	for word, level := range m {
		if level < 1 || level > 4 {
			t.Errorf("NGSL[%q] = %d; want 1-4", word, level)
		}
	}
}

func TestLoadNGSLAllLevelsPresent(t *testing.T) {
	m := LoadNGSL()
	counts := make(map[int]int)
	for _, level := range m {
		counts[level]++
	}
	for level := 1; level <= 4; level++ {
		if counts[level] == 0 {
			t.Errorf("no words at NGSL level %d", level)
		}
	}
	t.Logf("NGSL level distribution: %v", counts)
}

func TestLoadAWL(t *testing.T) {
	m := LoadAWL()
	if len(m) == 0 {
		t.Fatal("LoadAWL() returned empty map")
	}
	if len(m) < 2500 {
		t.Errorf("LoadAWL() has %d entries; want at least 2500", len(m))
	}
	if got, ok := m["analysis"]; !ok {
		t.Error("word \"analysis\" not found in AWL map")
	} else if got != 4 {
		t.Errorf("AWL[\"analysis\"] = %d; want 4", got)
	}
	if got, ok := m["confirm"]; !ok {
		t.Error("word \"confirm\" not found in AWL map")
	} else if got != 5 {
		t.Errorf("AWL[\"confirm\"] = %d; want 5", got)
	}
}

func TestLoadAWLLevelRange(t *testing.T) {
	m := LoadAWL()
	for word, level := range m {
		if level < 4 || level > 5 {
			t.Errorf("AWL[%q] = %d; want 4-5", word, level)
		}
	}
}

func TestLoadIrregulars(t *testing.T) {
	m := LoadIrregulars()
	if len(m) == 0 {
		t.Fatal("LoadIrregulars() returned empty map")
	}
	if len(m) < 300 {
		t.Errorf("LoadIrregulars() has %d entries; want at least 300", len(m))
	}
	tests := []struct {
		variant string
		lemma   string
	}{
		{"went", "go"},
		{"gone", "go"},
		{"children", "child"},
		{"better", "good"},
		{"best", "good"},
		{"saw", "see"},
		{"seen", "see"},
		{"women", "woman"},
		{"mice", "mouse"},
		{"thought", "think"},
		{"written", "write"},
	}
	for _, tt := range tests {
		t.Run(tt.variant, func(t *testing.T) {
			got, ok := m[tt.variant]
			if !ok {
				t.Errorf("variant %q not found in irregulars map", tt.variant)
				return
			}
			if got != tt.lemma {
				t.Errorf("Irregulars[%q] = %q; want %q", tt.variant, got, tt.lemma)
			}
		})
	}
}

func TestLoadIrregularPastParticiples(t *testing.T) {
	m := LoadIrregularPastParticiples()
	if len(m) == 0 {
		t.Fatal("LoadIrregularPastParticiples() returned empty map")
	}
	ppWords := []string{"gone", "seen", "written", "broken", "taken", "given", "chosen", "driven", "eaten", "spoken", "frozen", "stolen"}
	for _, w := range ppWords {
		if !m[w] {
			t.Errorf("expected %q to be in past participles set", w)
		}
	}
	if m["went"] {
		t.Error("\"went\" should not be in past participles set (it's past tense only)")
	}
}

func TestLoadStopwords(t *testing.T) {
	m := LoadStopwords()
	if len(m) == 0 {
		t.Fatal("LoadStopwords() returned empty map")
	}
	if len(m) < 150 {
		t.Errorf("LoadStopwords() has %d entries; want at least 150", len(m))
	}
	expected := []string{"the", "a", "an", "is", "are", "was", "were", "i", "you", "he", "she", "it", "we", "they", "in", "on", "at", "to", "for", "with", "and", "but", "or", "not"}
	for _, w := range expected {
		if !m[w] {
			t.Errorf("stopword %q not found", w)
		}
	}
	if !m["i"] {
		t.Error("pronoun \"i\" (lowercase) must be in stopwords set")
	}
}

func TestLoadAbbreviations(t *testing.T) {
	m := LoadAbbreviations()
	if len(m) == 0 {
		t.Fatal("LoadAbbreviations() returned empty map")
	}
	if len(m) < 40 {
		t.Errorf("LoadAbbreviations() has %d entries; want at least 40", len(m))
	}
	expected := []string{"mr.", "mrs.", "ms.", "dr.", "prof.", "etc.", "vs.", "i.e.", "e.g.", "jan.", "feb."}
	for _, w := range expected {
		if !m[w] {
			t.Errorf("abbreviation %q not found", w)
		}
	}
}

func TestLoadIdempotent(t *testing.T) {
	o1 := LoadOxford()
	o2 := LoadOxford()
	if len(o1) != len(o2) {
		t.Error("LoadOxford() returned different maps on second call")
	}
	n1 := LoadNGSL()
	n2 := LoadNGSL()
	if len(n1) != len(n2) {
		t.Error("LoadNGSL() returned different maps on second call")
	}
}

func TestLoadersReturnDefensiveCopies(t *testing.T) {
	o1 := LoadOxford()
	original := o1["the"]
	o1["the"] = 99
	o2 := LoadOxford()
	if got := o2["the"]; got != original {
		t.Fatalf("LoadOxford() exposed mutable internal map; got %d, want %d", got, original)
	}

	n1 := LoadNGSL()
	n1["the"] = 99
	n2 := LoadNGSL()
	if got := n2["the"]; got == 99 {
		t.Fatal("LoadNGSL() exposed mutable internal map")
	}

	i1 := LoadIrregulars()
	i1["went"] = "mutated"
	i2 := LoadIrregulars()
	if got := i2["went"]; got == "mutated" {
		t.Fatal("LoadIrregulars() exposed mutable internal map")
	}

	s1 := LoadStopwords()
	s1["the"] = false
	s2 := LoadStopwords()
	if !s2["the"] {
		t.Fatal("LoadStopwords() exposed mutable internal map")
	}
}

func TestParseCSVFailuresReturnError(t *testing.T) {
	badCSV := []byte("\"unterminated")
	if _, err := parseOxfordCSV(badCSV); err == nil {
		t.Fatal("parseOxfordCSV() expected error for malformed CSV")
	}
	if _, err := parseNGSLCVS(badCSV); err == nil {
		t.Fatal("parseNGSLCVS() expected error for malformed CSV")
	}
	if _, err := parseAWLCSV(badCSV); err == nil {
		t.Fatal("parseAWLCSV() expected error for malformed CSV")
	}
	if _, _, err := parseIrregularsCSV(badCSV); err == nil {
		t.Fatal("parseIrregularsCSV() expected error for malformed CSV")
	}
}

func TestWordLookupAPIs(t *testing.T) {
	if got, ok := OxfordLevel("book"); !ok || got != 1 {
		t.Fatalf("OxfordLevel(\"book\") = (%d, %v); want (1, true)", got, ok)
	}

	if got, ok := NGSLLevel("the"); !ok || got != 1 {
		t.Fatalf("NGSLLevel(\"the\") = (%d, %v); want (1, true)", got, ok)
	}

	if got, ok := AWLLevel("analysis"); !ok || got != 4 {
		t.Fatalf("AWLLevel(\"analysis\") = (%d, %v); want (4, true)", got, ok)
	}

	if got, ok := IrregularLemma("went"); !ok || got != "go" {
		t.Fatalf("IrregularLemma(\"went\") = (%q, %v); want (\"go\", true)", got, ok)
	}

	if _, ok := OxfordLevel("this-word-should-not-exist"); ok {
		t.Fatal("OxfordLevel() unexpectedly found missing word")
	}
}
