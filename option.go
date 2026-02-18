package cefr

import "math"

// config holds internal parameters for assessment.
type config struct {
	vocabWeight       float64 // default 0.50
	syntaxWeight      float64 // default 0.30
	readabilityWeight float64 // default 0.20
	fullAnalysis      bool    // default false
	samplingThreshold int     // default 10000
}

// Option configures the behavior of the Assess function.
type Option func(*config)

// defaultConfig returns a config with default values.
func defaultConfig() config {
	return config{
		vocabWeight:       0.50,
		syntaxWeight:      0.30,
		readabilityWeight: 0.20,
		fullAnalysis:      false,
		samplingThreshold: 10000,
	}
}

// WithWeights sets custom weights for vocabulary, syntax and readability
// sub-scores. The three values must sum to 1.0 (tolerance ±0.001).
// If the constraint is not satisfied, the option is silently ignored and
// default weights are kept.
func WithWeights(vocab, syntax, readability float64) Option {
	return func(c *config) {
		if vocab < 0 || syntax < 0 || readability < 0 {
			return
		}
		if math.Abs(vocab+syntax+readability-1.0) > 0.001 {
			return
		}
		c.vocabWeight = vocab
		c.syntaxWeight = syntax
		c.readabilityWeight = readability
	}
}

// WithFullAnalysis forces full-text analysis instead of sampling for long
// texts.
func WithFullAnalysis() Option {
	return func(c *config) {
		c.fullAnalysis = true
	}
}

// WithSamplingThreshold sets a custom word-count threshold above which
// sampling is applied. n must be > 0; otherwise the option is ignored.
func WithSamplingThreshold(n int) Option {
	return func(c *config) {
		if n > 0 {
			c.samplingThreshold = n
		}
	}
}
