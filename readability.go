package cefr

// calcFKGL computes the Flesch-Kincaid Grade Level.
//
//	FKGL = 0.39×(words/sentences) + 11.8×(syllables/words) - 15.59
func calcFKGL(wordCount, sentenceCount, syllableCount int) float64 {
	w := float64(wordCount)
	s := float64(sentenceCount)
	syl := float64(syllableCount)
	return 0.39*(w/s) + 11.8*(syl/w) - 15.59
}

// calcFRE computes the Flesch Reading Ease score.
//
//	FRE = 206.835 - 1.015×(words/sentences) - 84.6×(syllables/words)
func calcFRE(wordCount, sentenceCount, syllableCount int) float64 {
	w := float64(wordCount)
	s := float64(sentenceCount)
	syl := float64(syllableCount)
	return 206.835 - 1.015*(w/s) - 84.6*(syl/w)
}

// calcCLI computes the Coleman-Liau Index.
//
//	L = (charCount / wordCount) × 100
//	S = (sentenceCount / wordCount) × 100
//	CLI = 0.0588×L - 0.296×S - 15.8
func calcCLI(charCount, wordCount, sentenceCount int) float64 {
	w := float64(wordCount)
	l := (float64(charCount) / w) * 100.0
	s := (float64(sentenceCount) / w) * 100.0
	return 0.0588*l - 0.296*s - 15.8
}

// gradeAnchors maps grade-level values to CEFR scores (1.0–6.0).
// Used by both fkglToScore and cliToScore.
var gradeAnchors = [][2]float64{
	{1.5, 1.0},
	{3.5, 2.0},
	{6.0, 3.0},
	{9.0, 4.0},
	{12.0, 5.0},
	{14.0, 6.0},
}

// interpolateScore performs piecewise linear interpolation on the given anchors.
// Each anchor is {inputValue, score}. Input values must be in ascending order.
// Values below the first anchor clamp to the first score; above the last clamp
// to the last score.
func interpolateScore(anchors [][2]float64, value float64) float64 {
	if value <= anchors[0][0] {
		return anchors[0][1]
	}
	for i := 1; i < len(anchors); i++ {
		if value <= anchors[i][0] {
			x0, y0 := anchors[i-1][0], anchors[i-1][1]
			x1, y1 := anchors[i][0], anchors[i][1]
			return y0 + (value-x0)/(x1-x0)*(y1-y0)
		}
	}
	return anchors[len(anchors)-1][1]
}

// fkglToScore maps a Flesch-Kincaid Grade Level to a CEFR score (1.0–6.0).
func fkglToScore(fkgl float64) float64 {
	return interpolateScore(gradeAnchors, fkgl)
}

// cliToScore maps a Coleman-Liau Index to a CEFR score (1.0–6.0).
// CLI approximates US grade level, so the same mapping as FKGL is used.
func cliToScore(cli float64) float64 {
	return interpolateScore(gradeAnchors, cli)
}

// freAnchors maps Flesch Reading Ease values to CEFR scores.
// FRE is inverse: higher FRE = easier text = lower CEFR.
// Anchors are in descending FRE order for clarity, but interpolateScore
// expects ascending input, so we store them reversed (ascending FRE).
var freAnchors = [][2]float64{
	{45.0, 6.0},
	{55.0, 5.0},
	{65.0, 4.0},
	{75.0, 3.0},
	{85.0, 2.0},
	{95.0, 1.0},
}

// freToScore maps a Flesch Reading Ease score to a CEFR score (1.0–6.0).
func freToScore(fre float64) float64 {
	// FRE anchors are ascending in FRE value but descending in score.
	// Clamp at boundaries.
	if fre >= freAnchors[len(freAnchors)-1][0] {
		return freAnchors[len(freAnchors)-1][1]
	}
	if fre <= freAnchors[0][0] {
		return freAnchors[0][1]
	}
	for i := 1; i < len(freAnchors); i++ {
		if fre <= freAnchors[i][0] {
			x0, y0 := freAnchors[i-1][0], freAnchors[i-1][1]
			x1, y1 := freAnchors[i][0], freAnchors[i][1]
			return y0 + (fre-x0)/(x1-x0)*(y1-y0)
		}
	}
	return freAnchors[len(freAnchors)-1][1]
}

// analyzeReadability computes readability metrics and a combined CEFR score.
//
// Weights: 50% FKGL + 30% FRE + 20% CLI.
// Returns a default ReadabilityResult with Score 1.0 if wordCount or
// sentenceCount is zero (avoids division by zero).
func analyzeReadability(wordCount, sentenceCount, syllableCount, charCount int) ReadabilityResult {
	if wordCount == 0 || sentenceCount == 0 {
		return ReadabilityResult{Score: 1.0}
	}

	fkgl := calcFKGL(wordCount, sentenceCount, syllableCount)
	fre := calcFRE(wordCount, sentenceCount, syllableCount)
	cli := calcCLI(charCount, wordCount, sentenceCount)

	score := 0.50*fkglToScore(fkgl) + 0.30*freToScore(fre) + 0.20*cliToScore(cli)

	return ReadabilityResult{
		Score: score,
		FKGL:  fkgl,
		FRE:   fre,
		CLI:   cli,
	}
}
