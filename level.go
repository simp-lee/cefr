package cefr

// Level represents a CEFR proficiency level as an integer.
type Level int

// CEFR level constants.
const (
	A1 Level = 1
	A2 Level = 2
	B1 Level = 3
	B2 Level = 4
	C1 Level = 5
	C2 Level = 6
)

// String returns the human-readable CEFR label for the level.
func (l Level) String() string {
	switch l {
	case A1:
		return "A1"
	case A2:
		return "A2"
	case B1:
		return "B1"
	case B2:
		return "B2"
	case C1:
		return "C1"
	case C2:
		return "C2"
	default:
		return "Unknown"
	}
}

// scoreToLevel maps a continuous score to a CEFR level label using
// left-closed, right-open intervals:
//
//	[1.0, 1.5) → A1
//	[1.5, 2.5) → A2
//	[2.5, 3.5) → B1
//	[3.5, 4.5) → B2
//	[4.5, 5.5) → C1
//	[5.5, 6.0] → C2
func scoreToLevel(score float64) string {
	score = clampScore(score)
	switch {
	case score < 1.5:
		return A1.String()
	case score < 2.5:
		return A2.String()
	case score < 3.5:
		return B1.String()
	case score < 4.5:
		return B2.String()
	case score < 5.5:
		return C1.String()
	default:
		return C2.String()
	}
}

// levelToBaseScore returns the base score for a given integer level value.
// Out-of-range values are clamped to [1.0, 6.0].
func levelToBaseScore(level int) float64 {
	if level < int(A1) {
		return float64(A1)
	}
	if level > int(C2) {
		return float64(C2)
	}
	return float64(level)
}

// clampScore constrains a score to the valid range [1.0, 6.0].
func clampScore(score float64) float64 {
	if score < 1.0 {
		return 1.0
	}
	if score > 6.0 {
		return 6.0
	}
	return score
}
