package readability

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var wordsCountTest = map[string]float64{
	"creature some a": 3,
	"some a":          2,
	"a":               1,
}

func TestWords(t *testing.T) {
	for word, expectedWordsCount := range wordsCountTest {
		_, count := Words(word)
		assert.Equal(t, expectedWordsCount, count)
	}
}

var sentenceCountTest = map[string]float64{
	"A set of words that is complete. In itself, typically containing a subject.": 2,
	"A set of words that is complete.":                                            1,
	"A set of words":                                                              1,
}

func TestSentenceCount(t *testing.T) {
	for sentence, expectedSentenceCount := range sentenceCountTest {
		_, count := Sentences(sentence)
		assert.Equal(t, expectedSentenceCount, count)
	}
}

var characterCountTests = []struct {
	input    string
	expected float64
}{
	{"aoms", 4},
	{"šeimeną", 7},
	{"русский язык", 11},
	{"добавление", 10},
}

func TestCharacterCount(t *testing.T) {
	for _, test := range characterCountTests {
		words, _ := Words(test.input)
		got := CharactersCount(words)
		assert.Equal(t, test.expected, got)
	}
}
