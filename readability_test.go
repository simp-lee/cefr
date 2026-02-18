package cefr

import (
	"math"
	"testing"
)

func almostEqual(a, b, tol float64) bool {
	return math.Abs(a-b) < tol
}

func TestCalcFKGL(t *testing.T) {
	tests := []struct {
		name         string
		words, sents int
		syllables    int
		want         float64
	}{
		// 100 words, 5 sentences, 150 syllables
		// 0.39*(100/5) + 11.8*(150/100) - 15.59 = 0.39*20 + 11.8*1.5 - 15.59 = 7.8 + 17.7 - 15.59 = 9.91
		{"typical", 100, 5, 150, 9.91},
		// 50 words, 5 sentences, 60 syllables
		// 0.39*10 + 11.8*1.2 - 15.59 = 3.9 + 14.16 - 15.59 = 2.47
		{"easy", 50, 5, 60, 2.47},
		// 100 words, 10 sentences, 140 syllables
		// 0.39*(100/10) + 11.8*(140/100) - 15.59 = 3.9 + 16.52 - 15.59 = 4.83
		{"reference", 100, 10, 140, 4.83},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcFKGL(tt.words, tt.sents, tt.syllables)
			if !almostEqual(got, tt.want, 0.01) {
				t.Errorf("calcFKGL(%d,%d,%d) = %.4f; want %.4f", tt.words, tt.sents, tt.syllables, got, tt.want)
			}
		})
	}
}

func TestCalcFRE(t *testing.T) {
	tests := []struct {
		name         string
		words, sents int
		syllables    int
		want         float64
	}{
		// 100 words, 5 sentences, 150 syllables
		// 206.835 - 1.015*20 - 84.6*1.5 = 206.835 - 20.3 - 126.9 = 59.635
		{"typical", 100, 5, 150, 59.635},
		// 100 words, 10 sentences, 140 syllables
		// 206.835 - 1.015*10 - 84.6*1.4 = 206.835 - 10.15 - 118.44 = 78.245
		{"reference", 100, 10, 140, 78.245},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcFRE(tt.words, tt.sents, tt.syllables)
			if !almostEqual(got, tt.want, 0.01) {
				t.Errorf("calcFRE(%d,%d,%d) = %.4f; want %.4f", tt.words, tt.sents, tt.syllables, got, tt.want)
			}
		})
	}
}

func TestCalcCLI(t *testing.T) {
	tests := []struct {
		name         string
		chars, words int
		sents        int
		want         float64
	}{
		// 400 chars, 100 words, 5 sentences
		// L = (400/100)*100 = 400, S = (5/100)*100 = 5
		// 0.0588*400 - 0.296*5 - 15.8 = 23.52 - 1.48 - 15.8 = 6.24
		{"typical", 400, 100, 5, 6.24},
		// 350 chars, 100 words, 10 sentences
		// L = (350/100)*100 = 350, S = (10/100)*100 = 10
		// 0.0588*350 - 0.296*10 - 15.8 = 20.58 - 2.96 - 15.8 = 1.82
		{"reference", 350, 100, 10, 1.82},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcCLI(tt.chars, tt.words, tt.sents)
			if !almostEqual(got, tt.want, 0.01) {
				t.Errorf("calcCLI(%d,%d,%d) = %.4f; want %.4f", tt.chars, tt.words, tt.sents, got, tt.want)
			}
		})
	}
}

func TestFkglToScore(t *testing.T) {
	tests := []struct {
		name string
		fkgl float64
		want float64
	}{
		// Anchor points
		{"grade_1.5", 1.5, 1.0},
		{"grade_3.5", 3.5, 2.0},
		{"grade_6.0", 6.0, 3.0},
		{"grade_9.0", 9.0, 4.0},
		{"grade_12.0", 12.0, 5.0},
		{"grade_14.0", 14.0, 6.0},
		// Interpolation: midpoint between 1.5→1.0 and 3.5→2.0 is grade 2.5→1.5
		{"grade_2.5", 2.5, 1.5},
		// Below minimum → 1.0
		{"grade_0", 0.0, 1.0},
		// Above maximum → 6.0
		{"grade_20", 20.0, 6.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fkglToScore(tt.fkgl)
			if !almostEqual(got, tt.want, 0.01) {
				t.Errorf("fkglToScore(%.2f) = %.4f; want %.4f", tt.fkgl, got, tt.want)
			}
		})
	}
}

func TestFreToScore(t *testing.T) {
	tests := []struct {
		name string
		fre  float64
		want float64
	}{
		// Anchor points
		{"fre_95", 95.0, 1.0},
		{"fre_85", 85.0, 2.0},
		{"fre_75", 75.0, 3.0},
		{"fre_65", 65.0, 4.0},
		{"fre_55", 55.0, 5.0},
		{"fre_45", 45.0, 6.0},
		// Interpolation: midpoint between 95→1.0 and 85→2.0 is FRE 90→1.5
		{"fre_90", 90.0, 1.5},
		// Above 95 → 1.0
		{"fre_100", 100.0, 1.0},
		// Below 45 → 6.0
		{"fre_30", 30.0, 6.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := freToScore(tt.fre)
			if !almostEqual(got, tt.want, 0.01) {
				t.Errorf("freToScore(%.2f) = %.4f; want %.4f", tt.fre, got, tt.want)
			}
		})
	}
}

func TestCliToScore(t *testing.T) {
	// CLI uses same mapping as FKGL
	tests := []struct {
		name string
		cli  float64
		want float64
	}{
		{"cli_1.5", 1.5, 1.0},
		{"cli_3.5", 3.5, 2.0},
		{"cli_6.0", 6.0, 3.0},
		{"cli_9.0", 9.0, 4.0},
		{"cli_12.0", 12.0, 5.0},
		{"cli_14.0", 14.0, 6.0},
		{"cli_below_min", 0.0, 1.0},
		{"cli_above_max", 20.0, 6.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cliToScore(tt.cli)
			if !almostEqual(got, tt.want, 0.01) {
				t.Errorf("cliToScore(%.2f) = %.4f; want %.4f", tt.cli, got, tt.want)
			}
		})
	}
}

func TestAnalyzeReadability(t *testing.T) {
	t.Run("zero_words", func(t *testing.T) {
		r := analyzeReadability(0, 5, 0, 0)
		if r.Score != 1.0 {
			t.Errorf("zero words: Score = %.4f; want 1.0", r.Score)
		}
	})

	t.Run("zero_sentences", func(t *testing.T) {
		r := analyzeReadability(10, 0, 15, 40)
		if r.Score != 1.0 {
			t.Errorf("zero sentences: Score = %.4f; want 1.0", r.Score)
		}
	})

	t.Run("normal_text", func(t *testing.T) {
		// 100 words, 5 sentences, 150 syllables, 400 chars
		r := analyzeReadability(100, 5, 150, 400)

		// FKGL = 9.91, fkglToScore(9.91) → interpolation between 9.0→4.0 and 12.0→5.0
		// (9.91-9.0)/(12.0-9.0) * (5.0-4.0) + 4.0 = 0.91/3.0 + 4.0 ≈ 4.3033
		expFKGLScore := 4.3033

		// FRE = 59.635, freToScore(59.635) → between 65→4.0 and 55→5.0
		// (65-59.635)/(65-55) * (5.0-4.0) + 4.0 = 5.365/10 + 4.0 ≈ 4.5365
		expFREScore := 4.5365

		// CLI = 6.24, cliToScore(6.24) → between 6.0→3.0 and 9.0→4.0
		// (6.24-6.0)/(9.0-6.0) * (4.0-3.0) + 3.0 = 0.24/3.0 + 3.0 ≈ 3.08
		expCLIScore := 3.08

		expScore := 0.50*expFKGLScore + 0.30*expFREScore + 0.20*expCLIScore

		if !almostEqual(r.FKGL, 9.91, 0.01) {
			t.Errorf("FKGL = %.4f; want 9.91", r.FKGL)
		}
		if !almostEqual(r.FRE, 59.635, 0.01) {
			t.Errorf("FRE = %.4f; want 59.635", r.FRE)
		}
		if !almostEqual(r.CLI, 6.24, 0.01) {
			t.Errorf("CLI = %.4f; want 6.24", r.CLI)
		}
		if !almostEqual(r.Score, expScore, 0.01) {
			t.Errorf("Score = %.4f; want %.4f", r.Score, expScore)
		}
	})

	t.Run("reference_100w_10s_140syl_350c", func(t *testing.T) {
		// 100 words, 10 sentences, 140 syllables, 350 chars
		r := analyzeReadability(100, 10, 140, 350)

		// FKGL = 0.39*10 + 11.8*1.4 - 15.59 = 4.83
		// FRE = 206.835 - 1.015*10 - 84.6*1.4 = 78.245
		// CLI: L=350, S=10 → 0.0588*350 - 0.296*10 - 15.8 = 1.82
		if !almostEqual(r.FKGL, 4.83, 0.01) {
			t.Errorf("FKGL = %.4f; want 4.83", r.FKGL)
		}
		if !almostEqual(r.FRE, 78.245, 0.01) {
			t.Errorf("FRE = %.4f; want 78.245", r.FRE)
		}
		if !almostEqual(r.CLI, 1.82, 0.01) {
			t.Errorf("CLI = %.4f; want 1.82", r.CLI)
		}

		// Verify composite: 0.50*fkglToScore(4.83) + 0.30*freToScore(78.245) + 0.20*cliToScore(1.82)
		expScore := 0.50*fkglToScore(4.83) + 0.30*freToScore(78.245) + 0.20*cliToScore(1.82)
		if !almostEqual(r.Score, expScore, 0.02) {
			t.Errorf("Score = %.4f; want %.4f", r.Score, expScore)
		}
	})
}
