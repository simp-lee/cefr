# cefr

A Pure Go CEFR Text Difficulty Assessment Library.

Assess the difficulty level of any English text according to the [Common European Framework of Reference for Languages (CEFR)](https://www.coe.int/en/web/common-european-framework-reference-languages) — from A1 (Beginner) to C2 (Proficiency).

## Features

- **Multi-feature fusion algorithm** — vocabulary (50%) + syntax (30%) + readability (20%)
- **Zero external dependencies** — built entirely on the Go standard library
- **Embedded word lists** — Oxford 5000, NGSL, AWL compiled into the binary
- **CEFR A1–C2 output** — continuous score (1.0–6.0) mapped to a discrete level
- **Confidence scoring** — reliability estimate based on text length and segment consistency
- **Long-text sampling** — automatic three-segment sampling for texts over 10 000 words
- **Configurable weights** — functional options pattern for customisation
- **Goroutine-safe** — no shared mutable state; safe for concurrent use

## Installation

```bash
go get github.com/simp-lee/cefr
```

Requires **Go 1.25** or later.

## Quick Start

### Basic Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/simp-lee/cefr"
)

func main() {
	text := "The cat sat on the mat."
	result, err := cefr.Assess(text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Level: %s, Score: %.2f, Confidence: %.2f\n",
		result.Level, result.Score, result.Confidence)
}
```

### Custom Weights

```go
result, err := cefr.Assess(text,
	cefr.WithWeights(0.60, 0.25, 0.15), // vocab 60%, syntax 25%, readability 15%
)
```

### Other Options

```go
// Force full analysis (disable sampling for long texts)
result, err := cefr.Assess(text, cefr.WithFullAnalysis())

// Set a custom sampling threshold (default: 10 000 words)
result, err := cefr.Assess(text, cefr.WithSamplingThreshold(5000))
```

## Result

The `Assess` function returns a `Result` struct with the following fields:

| Field | Type | Description |
|-------|------|-------------|
| `Level` | `string` | CEFR level label: `"A1"` through `"C2"` |
| `Score` | `float64` | Continuous score from 1.0 to 6.0 |
| `Confidence` | `float64` | Assessment reliability from 0.0 to 1.0 |
| `Vocab` | `VocabResult` | Vocabulary analysis details |
| `Syntax` | `SyntaxResult` | Syntactic complexity details |
| `Readability` | `ReadabilityResult` | Readability formula details |
| `WordCount` | `int` | Total word count |
| `SentenceCount` | `int` | Total sentence count |

### VocabResult

| Field | Type | Description |
|-------|------|-------------|
| `Score` | `float64` | Vocabulary sub-score (1.0–6.0) |
| `Distribution` | `map[string]float64` | CEFR level distribution of content words |
| `UnknownRatio` | `float64` | Ratio of words not found in any word list |
| `ContentWords` | `int` | Number of content words analyzed |

### SyntaxResult

| Field | Type | Description |
|-------|------|-------------|
| `Score` | `float64` | Syntax sub-score (1.0–6.0) |
| `AvgSentenceLength` | `float64` | Average words per sentence |
| `SubordinationIndex` | `float64` | Ratio of subordinate clauses |
| `PassiveRate` | `float64` | Ratio of passive voice constructions |
| `ConnectorDiversity` | `int` | Number of distinct connector types used |

### ReadabilityResult

| Field | Type | Description |
|-------|------|-------------|
| `Score` | `float64` | Readability sub-score (1.0–6.0) |
| `FKGL` | `float64` | Flesch-Kincaid Grade Level |
| `FRE` | `float64` | Flesch Reading Ease |
| `CLI` | `float64` | Coleman-Liau Index |

## CEFR Level Mapping

| Score Range | Level | Description |
|-------------|-------|-------------|
| [1.0, 1.5) | A1 | Beginner |
| [1.5, 2.5) | A2 | Elementary |
| [2.5, 3.5) | B1 | Intermediate |
| [3.5, 4.5) | B2 | Upper Intermediate |
| [4.5, 5.5) | C1 | Advanced |
| [5.5, 6.0] | C2 | Proficiency |

## Algorithm Overview

### Multi-Feature Fusion

The final score is a weighted sum of three dimensions:

```
Score = Vocab × 0.50 + Syntax × 0.30 + Readability × 0.20
```

### Vocabulary Analysis

- Looks up each content word against a layered word-list hierarchy: **Oxford 5000 → AWL → NGSL**
- Uses **heuristic lemmatisation** (suffix stripping + irregular-form lookup) to maximise list coverage
- Computes the CEFR distribution of matched words and derives a score from the **P80 percentile**

### Syntax Analysis

- **Average Sentence Length (ASL)** — mapped to CEFR via piecewise interpolation
- **Subordination Index** — ratio of subordinating conjunctions to sentences
- **Passive Voice Rate** — heuristic detection of passive constructions
- **Connector Diversity** — count of distinct discourse connector types

### Readability Formulas

Three classic readability measures, each interpolated to a 1.0–6.0 CEFR scale:

- **Flesch-Kincaid Grade Level (FKGL)**
- **Flesch Reading Ease (FRE)**
- **Coleman-Liau Index (CLI)**

### Long-Text Sampling

For texts exceeding 10 000 words (configurable), three ~1 000-word segments are extracted from the **beginning**, **middle**, and **end** of the text. Results are merged and variance is used to adjust the confidence score.

## Benchmarks

Measured on an Intel Core i5-4590 @ 3.30 GHz (4 cores), Go 1.25, Linux amd64:

| Text Size | Time | Allocs | Memory |
|-----------|------|--------|--------|
| 100 words | ~217 µs | 596 | 48 KB |
| 1 000 words | ~1.9 ms | 5 176 | 410 KB |
| 10 000 words | ~8.1 ms | 19 985 | 1.7 MB |
| 50 000 words (sampled) | ~14 ms | 36 489 | 3.3 MB |
| 50 000 words (full) | ~83 ms | 253 255 | 25.7 MB |

Sampling keeps latency nearly constant regardless of input size.

## Limitations

- Uses **heuristic rules**, not NLP/ML models — results are best-effort approximations
- Passive voice detection may produce **false positives** (e.g. adjectives resembling past participles)
- Syllable counting accuracy is approximately **85–90%** (rule-based, no pronunciation dictionary)
- Assessment of **non-standard English** (slang, dialect, code-mixed text) may be inaccurate
- Expected accuracy is **±1 CEFR level** compared to human expert judgement

## License

[MIT](LICENSE)
