package dalechall

// Dale-Chall is one of the most accurate readability metrics. Rather than rely on syllable
// counts to identify diffult words, Dale-Chall incorporates a list of 3,000 easy words
// which were understood by 80% of fourth-grade students. The readability score is then
// computed based on how many words pressent in the passage are not in the list of easy words.

// A score of 4.9 or lower indicates the passage is easily readable by the average 4th grade.
// Scores between 9.0 and 9.9 indicate the passage is at a college level of readability.

import (
	"math"
	"regexp"
	"strings"

	"github.com/20k-ultra/readability"
)

// Calculate implements Dale-Chall Readability formula
func Calculate(text []byte) float32 {
	words := readability.GetWords(text, true)
	sentences := readability.GetSentences(text)
	ASL := readability.AverageSentenceLength(words, sentences)
	PDW := PercentageofDifficultWords(words) * 100
	output := formula(PDW, ASL)
	return float32(math.Round(float64(output)*100) / 100)
}

// PercentageofDifficultWords returns percentage of words that are considered difficult by Dale Chall formula
func PercentageofDifficultWords(words []string) float32 {
	wordCount := len(words)
	difficultWordCount := 0
	difficultWords := make(map[string]string)
	for _, word := range words {
		regex := regexp.MustCompile(`\d`)
		if regex.ReplaceAllString(word, "") == "" {
			continue
		}
		word = strings.ToLower(word)
		if _, has := wordList[word]; !has {
			if _, alreadyKnow := difficultWords[word]; !alreadyKnow {
				difficultWords[word] = ""
				difficultWordCount++
			}
		}
	}
	return float32(difficultWordCount) / float32(wordCount)
}

func formula(percentageOfDifficultWords float32, averageSentenceLength float32) float32 {
	value := 0.1579*(percentageOfDifficultWords) + 0.0496*averageSentenceLength
	if percentageOfDifficultWords > 5 {
		value += 3.6365 // adjust score
	}
	return value
}
