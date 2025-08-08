package fake

import (
	"strings"

	"github.com/go-faker/faker/v4"
)

// Words returns a string containing exactly n words joined by a single space.
// If n <= 0, it returns an empty string.
func Words(n int) string {
	if n <= 0 {
		return ""
	}
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = faker.Word()
	}
	return strings.Join(words, " ")
}

// Paragraph returns a paragraph composed of the given number of sentences,
// where each sentence contains exactly wordsPerSentence words and ends with a period.
// If sentences <= 0 or wordsPerSentence <= 0, it returns an empty string.
func Paragraph(sentences int, wordsPerSentence int) string {
	if sentences <= 0 || wordsPerSentence <= 0 {
		return ""
	}
	var b strings.Builder
	for s := 0; s < sentences; s++ {
		if s > 0 {
			b.WriteByte(' ')
		}
		// Build a sentence
		for w := 0; w < wordsPerSentence; w++ {
			if w > 0 {
				b.WriteByte(' ')
			}
			b.WriteString(faker.Word())
		}
		b.WriteByte('.')
	}
	return b.String()
}
