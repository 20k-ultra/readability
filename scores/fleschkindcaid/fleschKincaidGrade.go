package fleschkindcaid

// The Flesch-Kincaid Grade Level heuristic indicates that the text
// can be read by the average student in the specified grade level.

import (
	"math"

	"github.com/20k-ultra/readability"
)

// Calculate implements Flesch-Kincaid Grade Level formula
func Calculate(text []byte) float32 {
	words := readability.GetWords(text, true)
	sentences := readability.GetSentences(text)
	ASL := readability.AverageSentenceLength(words, sentences)
	ASW := readability.CountSyllablesPerWord(words)
	output := formula(ASL, ASW)
	return float32(math.Round(float64(output)*100) / 100)
}

func formula(averageSentenceLength float32, AvgSyllablesPerWord float32) float32 {
	return (0.39 * averageSentenceLength) + (11.8 * AvgSyllablesPerWord) - 15.59
}
