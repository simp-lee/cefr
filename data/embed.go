// Package data holds embedded word frequency and linguistic reference data
// used by the cefr package for vocabulary lookup and analysis.
package data

import (
	"bufio"
	_ "embed"
	"encoding/csv"
	"fmt"
	"maps"
	"strconv"
	"strings"
	"sync"
)

//go:embed oxford5000.csv
var oxfordCSV []byte

//go:embed ngsl.csv
var ngslCSV []byte

//go:embed awl.csv
var awlCSV []byte

//go:embed irregulars.csv
var irregularsCSV []byte

//go:embed stopwords.csv
var stopwordsCSV []byte

//go:embed abbreviations.csv
var abbreviationsCSV []byte

var (
	oxfordOnce        sync.Once
	ngslOnce          sync.Once
	awlOnce           sync.Once
	irregularsOnce    sync.Once
	irregularPPOnce   sync.Once
	stopwordsOnce     sync.Once
	abbreviationsOnce sync.Once

	oxfordMap        map[string]int
	ngslMap          map[string]int
	awlMap           map[string]int
	irregularsMap    map[string]string
	irregularPPMap   map[string]bool
	stopwordsMap     map[string]bool
	abbreviationsMap map[string]bool

	loadErrMu sync.Mutex
	loadErr   error
)

func setLoadErr(err error) {
	if err == nil {
		return
	}
	loadErrMu.Lock()
	defer loadErrMu.Unlock()
	if loadErr == nil {
		loadErr = err
	}
}

// InitError returns the first dataset initialization error encountered,
// or nil when all embedded datasets were loaded successfully.
func InitError() error {
	LoadOxford()
	LoadNGSL()
	LoadAWL()
	LoadIrregulars()
	LoadIrregularPastParticiples()
	LoadStopwords()
	LoadAbbreviations()

	loadErrMu.Lock()
	defer loadErrMu.Unlock()
	if loadErr == nil {
		return nil
	}
	return loadErr
}

func cloneMapInt(src map[string]int) map[string]int {
	dst := make(map[string]int, len(src))
	maps.Copy(dst, src)
	return dst
}

func cloneMapString(src map[string]string) map[string]string {
	dst := make(map[string]string, len(src))
	maps.Copy(dst, src)
	return dst
}

func cloneMapBool(src map[string]bool) map[string]bool {
	dst := make(map[string]bool, len(src))
	maps.Copy(dst, src)
	return dst
}

func parseOxfordCSV(content []byte) (map[string]int, error) {
	levelMap := map[string]int{
		"a1": 1, "a2": 2, "b1": 3, "b2": 4, "c1": 5,
	}
	result := make(map[string]int)
	reader := csv.NewReader(strings.NewReader(string(content)))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, record := range records {
		if len(record) < 2 {
			continue
		}
		word := strings.TrimSpace(strings.ToLower(record[0]))
		level := strings.TrimSpace(strings.ToLower(record[1]))
		if mappedLevel, ok := levelMap[level]; ok && word != "" {
			result[word] = mappedLevel
		}
	}
	return result, nil
}

func parseNGSLCVS(content []byte) (map[string]int, error) {
	result := make(map[string]int)
	reader := csv.NewReader(strings.NewReader(string(content)))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, record := range records {
		if len(record) < 2 {
			continue
		}
		rank, parseErr := strconv.Atoi(strings.TrimSpace(record[0]))
		if parseErr != nil {
			continue
		}
		word := strings.TrimSpace(strings.ToLower(record[1]))
		if word == "" {
			continue
		}
		var level int
		switch {
		case rank <= 500:
			level = 1
		case rank <= 1200:
			level = 2
		case rank <= 2000:
			level = 3
		default:
			level = 4
		}
		result[word] = level
	}
	return result, nil
}

func parseAWLCSV(content []byte) (map[string]int, error) {
	result := make(map[string]int)
	reader := csv.NewReader(strings.NewReader(string(content)))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, record := range records {
		if len(record) < 2 {
			continue
		}
		sublist, parseErr := strconv.Atoi(strings.TrimSpace(record[0]))
		if parseErr != nil {
			continue
		}
		word := strings.TrimSpace(strings.ToLower(record[1]))
		if word == "" {
			continue
		}
		var level int
		if sublist <= 5 {
			level = 4
		} else {
			level = 5
		}
		result[word] = level
	}
	return result, nil
}

func parseIrregularsCSV(content []byte) (map[string]string, map[string]bool, error) {
	irregulars := make(map[string]string)
	irregularPP := make(map[string]bool)
	reader := csv.NewReader(strings.NewReader(string(content)))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	for _, record := range records {
		if len(record) < 2 {
			continue
		}
		variant := strings.TrimSpace(strings.ToLower(record[0]))
		lemma := strings.TrimSpace(strings.ToLower(record[1]))
		if variant != "" && lemma != "" {
			irregulars[variant] = lemma
		}
		if len(record) >= 3 {
			typ := strings.TrimSpace(strings.ToLower(record[2]))
			if typ == "pp" && variant != "" {
				irregularPP[variant] = true
			}
		}
	}
	return irregulars, irregularPP, nil
}

func parseWordSetLines(content []byte) (map[string]bool, error) {
	result := make(map[string]bool)
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	for scanner.Scan() {
		word := strings.TrimSpace(strings.ToLower(scanner.Text()))
		if word != "" {
			result[word] = true
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func loadOxfordShared() map[string]int {
	oxfordOnce.Do(func() {
		parsed, err := parseOxfordCSV(oxfordCSV)
		if err != nil {
			setLoadErr(fmt.Errorf("failed to parse embedded CSV oxford5000.csv: %w", err))
			oxfordMap = make(map[string]int)
			return
		}
		oxfordMap = parsed
	})
	return oxfordMap
}

func loadNGSLShared() map[string]int {
	ngslOnce.Do(func() {
		parsed, err := parseNGSLCVS(ngslCSV)
		if err != nil {
			setLoadErr(fmt.Errorf("failed to parse embedded CSV ngsl.csv: %w", err))
			ngslMap = make(map[string]int)
			return
		}
		ngslMap = parsed
	})
	return ngslMap
}

func loadAWLShared() map[string]int {
	awlOnce.Do(func() {
		parsed, err := parseAWLCSV(awlCSV)
		if err != nil {
			setLoadErr(fmt.Errorf("failed to parse embedded CSV awl.csv: %w", err))
			awlMap = make(map[string]int)
			return
		}
		awlMap = parsed
	})
	return awlMap
}

func loadIrregularsShared() map[string]string {
	irregularsOnce.Do(func() {
		parsedIrregulars, _, err := parseIrregularsCSV(irregularsCSV)
		if err != nil {
			setLoadErr(fmt.Errorf("failed to parse embedded CSV irregulars.csv: %w", err))
			irregularsMap = make(map[string]string)
			return
		}
		irregularsMap = parsedIrregulars
	})
	return irregularsMap
}

// OxfordLevel returns the Oxford CEFR level for a single word, if present.
func OxfordLevel(word string) (int, bool) {
	level, ok := loadOxfordShared()[word]
	return level, ok
}

// NGSLLevel returns the NGSL level for a single word, if present.
func NGSLLevel(word string) (int, bool) {
	level, ok := loadNGSLShared()[word]
	return level, ok
}

// AWLLevel returns the AWL level for a single word, if present.
func AWLLevel(word string) (int, bool) {
	level, ok := loadAWLShared()[word]
	return level, ok
}

// IrregularLemma returns the lemma for an irregular form, if present.
func IrregularLemma(word string) (string, bool) {
	lemma, ok := loadIrregularsShared()[word]
	return lemma, ok
}

// LoadOxford returns a map of word → CEFR level (1–5, where a1=1, a2=2, b1=3, b2=4, c1=5).
// The returned map must not be modified by the caller.
func LoadOxford() map[string]int {
	return cloneMapInt(loadOxfordShared())
}

// LoadNGSL returns a map of word → level (1–4, mapped by frequency rank):
//
//	1–500 → 1, 501–1200 → 2, 1201–2000 → 3, 2001–2800 → 4
//
// The returned map must not be modified by the caller.
func LoadNGSL() map[string]int {
	return cloneMapInt(loadNGSLShared())
}

// LoadAWL returns a map of word → level (4–5, mapped by sublist):
//
//	sublist 1–5 → 4, sublist 6–10 → 5
//
// The returned map must not be modified by the caller.
func LoadAWL() map[string]int {
	return cloneMapInt(loadAWLShared())
}

// LoadIrregulars returns a map of variant → lemma for irregular word forms.
// The returned map must not be modified by the caller.
func LoadIrregulars() map[string]string {
	return cloneMapString(loadIrregularsShared())
}

// LoadIrregularPastParticiples returns a set of irregular past participle forms
// (type=pp entries from the irregulars data). Used for passive voice detection.
// The returned map must not be modified by the caller.
func LoadIrregularPastParticiples() map[string]bool {
	irregularPPOnce.Do(func() {
		_, parsedPP, err := parseIrregularsCSV(irregularsCSV)
		if err != nil {
			setLoadErr(fmt.Errorf("failed to parse embedded CSV irregulars.csv: %w", err))
			irregularPPMap = make(map[string]bool)
			return
		}
		irregularPPMap = parsedPP
	})
	return cloneMapBool(irregularPPMap)
}

// LoadStopwords returns a set of common English stopwords.
// The returned map must not be modified by the caller.
func LoadStopwords() map[string]bool {
	stopwordsOnce.Do(func() {
		parsed, err := parseWordSetLines(stopwordsCSV)
		if err != nil {
			setLoadErr(fmt.Errorf("failed to parse embedded data stopwords.csv: %w", err))
			stopwordsMap = make(map[string]bool)
			return
		}
		stopwordsMap = parsed
	})
	return cloneMapBool(stopwordsMap)
}

// LoadAbbreviations returns a set of common English abbreviations (e.g., "mr.", "dr.").
// The returned map must not be modified by the caller.
func LoadAbbreviations() map[string]bool {
	abbreviationsOnce.Do(func() {
		parsed, err := parseWordSetLines(abbreviationsCSV)
		if err != nil {
			setLoadErr(fmt.Errorf("failed to parse embedded data abbreviations.csv: %w", err))
			abbreviationsMap = make(map[string]bool)
			return
		}
		abbreviationsMap = parsed
	})
	return cloneMapBool(abbreviationsMap)
}
