package cefr

import (
	"math"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := defaultConfig()

	if cfg.vocabWeight != 0.50 {
		t.Errorf("vocabWeight = %v; want 0.50", cfg.vocabWeight)
	}
	if cfg.syntaxWeight != 0.30 {
		t.Errorf("syntaxWeight = %v; want 0.30", cfg.syntaxWeight)
	}
	if cfg.readabilityWeight != 0.20 {
		t.Errorf("readabilityWeight = %v; want 0.20", cfg.readabilityWeight)
	}
	if cfg.fullAnalysis {
		t.Error("fullAnalysis = true; want false")
	}
	if cfg.samplingThreshold != 10000 {
		t.Errorf("samplingThreshold = %d; want 10000", cfg.samplingThreshold)
	}
}

func TestWithWeights(t *testing.T) {
	tests := []struct {
		name        string
		v, s, r     float64
		wantV       float64
		wantS       float64
		wantR       float64
		wantDefault bool // expect defaults to remain
	}{
		{
			name: "valid weights",
			v:    0.60, s: 0.25, r: 0.15,
			wantV: 0.60, wantS: 0.25, wantR: 0.15,
		},
		{
			name: "valid equal weights",
			v:    1.0 / 3, s: 1.0 / 3, r: 1.0 / 3,
			wantV: 1.0 / 3, wantS: 1.0 / 3, wantR: 1.0 / 3,
		},
		{
			name: "within float tolerance",
			v:    0.500001, s: 0.300000, r: 0.200000,
			wantV: 0.500001, wantS: 0.300000, wantR: 0.200000,
		},
		{
			name: "sum too high",
			v:    0.50, s: 0.30, r: 0.25,
			wantDefault: true,
		},
		{
			name: "sum too low",
			v:    0.10, s: 0.10, r: 0.10,
			wantDefault: true,
		},
		{
			name: "negative weight",
			v:    -0.10, s: 0.60, r: 0.50,
			wantDefault: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := defaultConfig()
			opt := WithWeights(tt.v, tt.s, tt.r)
			opt(&cfg)

			if tt.wantDefault {
				def := defaultConfig()
				if cfg.vocabWeight != def.vocabWeight || cfg.syntaxWeight != def.syntaxWeight || cfg.readabilityWeight != def.readabilityWeight {
					t.Errorf("expected default weights, got vocab=%v syntax=%v readability=%v",
						cfg.vocabWeight, cfg.syntaxWeight, cfg.readabilityWeight)
				}
			} else {
				if math.Abs(cfg.vocabWeight-tt.wantV) > 1e-9 {
					t.Errorf("vocabWeight = %v; want %v", cfg.vocabWeight, tt.wantV)
				}
				if math.Abs(cfg.syntaxWeight-tt.wantS) > 1e-9 {
					t.Errorf("syntaxWeight = %v; want %v", cfg.syntaxWeight, tt.wantS)
				}
				if math.Abs(cfg.readabilityWeight-tt.wantR) > 1e-9 {
					t.Errorf("readabilityWeight = %v; want %v", cfg.readabilityWeight, tt.wantR)
				}
			}
		})
	}
}

func TestWithFullAnalysis(t *testing.T) {
	cfg := defaultConfig()
	if cfg.fullAnalysis {
		t.Fatal("precondition: fullAnalysis should be false")
	}

	opt := WithFullAnalysis()
	opt(&cfg)

	if !cfg.fullAnalysis {
		t.Error("fullAnalysis = false; want true")
	}
}

func TestWithSamplingThreshold(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		wantValue int
	}{
		{"positive value", 5000, 5000},
		{"value of 1", 1, 1},
		{"zero ignored", 0, 10000},
		{"negative ignored", -1, 10000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := defaultConfig()
			opt := WithSamplingThreshold(tt.n)
			opt(&cfg)

			if cfg.samplingThreshold != tt.wantValue {
				t.Errorf("samplingThreshold = %d; want %d", cfg.samplingThreshold, tt.wantValue)
			}
		})
	}
}
