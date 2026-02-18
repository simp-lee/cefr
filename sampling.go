package cefr

import "strings"

// shouldSample reports whether the text is long enough to benefit from
// sampling rather than full analysis.
func shouldSample(wordCount, threshold int) bool {
	return wordCount > threshold
}

// sampleText extracts three ~1000-word segments (beginning, middle, end) from
// a long text, splitting at sentence boundaries. If the text has fewer than
// 3000 words, it returns the entire text as a single-element slice.
func sampleText(text string, wordCount int) []string {
	if wordCount < 3000 {
		return []string{text}
	}

	sentences := splitSentences(text)
	if len(sentences) == 0 {
		return []string{text}
	}

	const targetWords = 1000

	uniqueSegments := func(candidates ...string) []string {
		seen := make(map[string]bool)
		segments := make([]string, 0, len(candidates))
		for _, seg := range candidates {
			trimmed := strings.TrimSpace(seg)
			if trimmed == "" || seen[trimmed] {
				continue
			}
			seen[trimmed] = true
			segments = append(segments, trimmed)
		}
		if len(segments) == 0 {
			return []string{text}
		}
		return segments
	}

	// Helper: collect sentences from index start until we have >= targetWords words.
	// Returns the collected text and the next sentence index.
	collectSegment := func(start int) (string, int) {
		var b strings.Builder
		wordsCollected := 0
		i := start
		for i < len(sentences) && wordsCollected < targetWords {
			if b.Len() > 0 {
				b.WriteByte(' ')
			}
			b.WriteString(sentences[i])
			wordsCollected += len(strings.Fields(sentences[i]))
			i++
		}
		return b.String(), i
	}

	// Segment 1: beginning
	seg1, seg1End := collectSegment(0)

	// Segment 3: ending (collect from the end backward)
	// Find the start index such that sentences from there to the end have ~targetWords.
	seg3Start := len(sentences)
	seg3Words := 0
	for seg3Start > 0 && seg3Words < targetWords {
		seg3Start--
		seg3Words += len(strings.Fields(sentences[seg3Start]))
	}
	// Ensure segment 3 doesn't overlap with segment 1
	if seg3Start < seg1End {
		seg3Start = seg1End
	}
	var seg3Builder strings.Builder
	for i := seg3Start; i < len(sentences); i++ {
		if seg3Builder.Len() > 0 {
			seg3Builder.WriteByte(' ')
		}
		seg3Builder.WriteString(sentences[i])
	}
	seg3 := seg3Builder.String()

	// Segment 2: middle
	// Find the midpoint of remaining sentences between seg1End and seg3Start.
	midSentences := seg3Start - seg1End
	if midSentences <= 0 {
		// Not enough sentences for a separate middle segment.
		return uniqueSegments(seg1, seg3)
	}

	midStart := seg1End + midSentences/2
	// Adjust midStart backward so we center ~targetWords around the midpoint.
	halfTarget := targetWords / 2
	adjustedStart := midStart
	adjustWords := 0
	for adjustedStart > seg1End && adjustWords < halfTarget {
		adjustedStart--
		adjustWords += len(strings.Fields(sentences[adjustedStart]))
	}

	// Collect from adjustedStart, up to targetWords, not exceeding seg3Start.
	var seg2Builder strings.Builder
	wordsCollected := 0
	i := adjustedStart
	for i < seg3Start && wordsCollected < targetWords {
		if seg2Builder.Len() > 0 {
			seg2Builder.WriteByte(' ')
		}
		seg2Builder.WriteString(sentences[i])
		wordsCollected += len(strings.Fields(sentences[i]))
		i++
	}
	seg2 := seg2Builder.String()

	// Safety: if any segment is empty, fall back to not enough room for middle.
	if seg2 == "" {
		return uniqueSegments(seg1, seg3)
	}

	return []string{seg1, seg2, seg3}
}

// mergeSampledResults combines analysis results from multiple text segments
// into a single Result. If only one result is provided, it is returned as-is.
// The supplied config is used for the fusion weights so that caller-provided
// weights (via WithWeights) are respected.
func mergeSampledResults(results []Result, cfg config) Result {
	if len(results) == 0 {
		return Result{}
	}
	if len(results) == 1 {
		return results[0]
	}

	n := float64(len(results))

	// Aggregate vocab
	var vocabScoreSum float64
	var unknownRatioSum float64
	var contentWordsSum int
	distSum := make(map[string]float64)
	for _, r := range results {
		vocabScoreSum += r.Vocab.Score
		unknownRatioSum += r.Vocab.UnknownRatio
		contentWordsSum += r.Vocab.ContentWords
		for k, v := range r.Vocab.Distribution {
			distSum[k] += v
		}
	}
	distAvg := make(map[string]float64, len(distSum))
	for k, v := range distSum {
		distAvg[k] = v / n
	}

	// Aggregate syntax
	var syntaxScoreSum float64
	var aslSum, subIdxSum, passiveSum float64
	var connDivSum int
	for _, r := range results {
		syntaxScoreSum += r.Syntax.Score
		aslSum += r.Syntax.AvgSentenceLength
		subIdxSum += r.Syntax.SubordinationIndex
		passiveSum += r.Syntax.PassiveRate
		connDivSum += r.Syntax.ConnectorDiversity
	}

	// Aggregate readability
	var readScoreSum float64
	var fkglSum, freSum, cliSum float64
	for _, r := range results {
		readScoreSum += r.Readability.Score
		fkglSum += r.Readability.FKGL
		freSum += r.Readability.FRE
		cliSum += r.Readability.CLI
	}

	// Aggregate word/sentence counts (sum, not average)
	var wordCountSum, sentCountSum int
	for _, r := range results {
		wordCountSum += r.WordCount
		sentCountSum += r.SentenceCount
	}

	vocabScore := vocabScoreSum / n
	syntaxScore := syntaxScoreSum / n
	readScore := readScoreSum / n

	// Recompute overall score with the caller's weights.
	score := cfg.vocabWeight*vocabScore + cfg.syntaxWeight*syntaxScore + cfg.readabilityWeight*readScore
	score = clampScore(score)

	return Result{
		Level:      scoreToLevel(score),
		Score:      score,
		Confidence: 0, // Placeholder: computed by Step 13
		Vocab: VocabResult{
			Score:        vocabScore,
			Distribution: distAvg,
			UnknownRatio: unknownRatioSum / n,
			ContentWords: contentWordsSum,
		},
		Syntax: SyntaxResult{
			Score:              syntaxScore,
			AvgSentenceLength:  aslSum / n,
			SubordinationIndex: subIdxSum / n,
			PassiveRate:        passiveSum / n,
			ConnectorDiversity: int(float64(connDivSum)/n + 0.5), // rounded average
		},
		Readability: ReadabilityResult{
			Score: readScore,
			FKGL:  fkglSum / n,
			FRE:   freSum / n,
			CLI:   cliSum / n,
		},
		WordCount:     wordCountSum,
		SentenceCount: sentCountSum,
	}
}

// sampledVariance computes the population variance of the Score field across
// multiple sampled Results. Used by Step 13 for confidence calculation.
func sampledVariance(results []Result) float64 {
	if len(results) <= 1 {
		return 0
	}
	n := float64(len(results))
	var sum float64
	for _, r := range results {
		sum += r.Score
	}
	mean := sum / n

	var sqDiffSum float64
	for _, r := range results {
		d := r.Score - mean
		sqDiffSum += d * d
	}
	return sqDiffSum / n
}
