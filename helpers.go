package readability

import (
	"unicode/utf8"

	"github.com/linkosmos/aggtext"
)

// Sentences - splits text into sentences
// 1. sentence must begin with capital letter
// 2. sentence might end with: dot, exclamation mark or question (. ! ?)
// 3. ignores trailing symbols defined in 2. point
func Sentences(text string) (sentences []string, count float64) {
	sentences = aggtext.Sentences(text)
	return sentences, float64(len(sentences))
}

// Words - splits text into words by removing white space and punctuation
func Words(text string) (words []string, total float64) {
	words = aggtext.Words(text)
	return words, float64(len(words))
}

// CharactersCount - counts characters utf8 characters (runes) in a
// given array of words use Words() function to split for highest precision
func CharactersCount(words []string) float64 {
	var total int
	for _, word := range words {
		total += utf8.RuneCountInString(word)
	}
	return float64(total)
}
