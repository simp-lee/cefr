package cefr

import (
	"strings"
	"testing"
)

// sentenceTemplates provides natural-looking English sentences of varying
// complexity for benchmark text generation.
var sentenceTemplates = []string{
	"The teacher explained the lesson to the students in a clear and simple way.",
	"She decided to walk home after finishing her work at the office.",
	"Many people enjoy reading books during the long winter evenings.",
	"The children were playing in the park while their parents watched from a bench.",
	"He asked his friend to help him move the heavy furniture upstairs.",
	"It is important to eat a balanced diet and exercise regularly.",
	"The weather was warm and sunny so they went to the beach.",
	"Although the exam was difficult most students managed to pass.",
	"The company announced that it would hire fifty new employees next month.",
	"She has been studying English for three years and can speak it fluently.",
	"The old bridge was replaced by a modern structure made of steel and concrete.",
	"We should consider the environmental impact before making any decisions.",
	"The scientist published a paper about the effects of climate change on wildlife.",
	"They traveled across the country by train and visited many small towns.",
	"The government introduced new regulations to improve air quality in major cities.",
	"Despite the heavy rain the football match continued without interruption.",
	"Learning a new language requires patience dedication and regular practice.",
	"The documentary explored the history of ancient civilizations in great detail.",
	"The restaurant serves a wide variety of dishes from different cultures.",
	"After careful analysis the committee recommended significant changes to the policy.",
}

// generateText creates a text string containing approximately wordCount words
// by cycling through sentence templates.
func generateText(wordCount int) string {
	var b strings.Builder
	// Pre-estimate capacity: average ~8 chars/word + space.
	b.Grow(wordCount * 9)

	words := 0
	for words < wordCount {
		s := sentenceTemplates[words%len(sentenceTemplates)]
		if b.Len() > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(s)
		words += len(strings.Fields(s))
	}
	return b.String()
}

func BenchmarkAssess100Words(b *testing.B) {
	text := generateText(100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Assess(text)
	}
}

func BenchmarkAssess1000Words(b *testing.B) {
	text := generateText(1000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Assess(text)
	}
}

func BenchmarkAssess10000Words(b *testing.B) {
	text := generateText(10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Assess(text)
	}
}

func BenchmarkAssess50000Words(b *testing.B) {
	text := generateText(50000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Default options: samplingThreshold=10000 triggers sampling automatically.
		_, _ = Assess(text)
	}
}

func BenchmarkAssess50000WordsFull(b *testing.B) {
	text := generateText(50000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Assess(text, WithFullAnalysis())
	}
}

func BenchmarkLookupWordLevel(b *testing.B) {
	words := []string{
		"teacher", "students", "balance", "decision", "environmental",
		"analysis", "policy", "language", "practice", "history",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = lookupWordLevel(words[i%len(words)])
	}
}

func BenchmarkLemmatize(b *testing.B) {
	words := []string{
		"running", "studied", "boxes", "stopped", "happiest",
		"quickly", "went", "children", "better", "making",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lemmatize(words[i%len(words)])
	}
}
