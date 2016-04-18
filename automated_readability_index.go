package readability

// Automated readability index
// https://en.wikipedia.org/wiki/Automated_readability_index
const (
	ARI1 = 4.71
	ARI2 = 0.5
	ARI3 = 21.43
)

// AutomatedReadabilityIndex - the automated readability index (ARI) is a
// readability test designed to gauge the understandability of a text.
func AutomatedReadabilityIndex(text string) float64 {
	words, wordsCount := Words(text)
	if wordsCount == 0 {
		return 0
	}

	_, sentencesCount := Sentences(text)
	return ARI1*(CharactersCount(words)/wordsCount) + ARI2*(wordsCount/sentencesCount) - ARI3
}
