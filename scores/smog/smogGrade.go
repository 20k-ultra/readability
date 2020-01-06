package smog

// The SMOG (Simple Measure of Gobbledygook) grade is commonly used in health care.
// The score represents the number of years of education needed to understand a passage of writing.

import (
	"math"

	"github.com/20k-ultra/readability"
)

// Calculate implements Flesch Reading Ease Readability formula
func Calculate(text []byte) float32 {
	words := readability.GetWords(text, true)
	sentences := readability.GetSentences(text)
	ASL := readability.AverageSentenceLength(words, sentences)
	ASW := readability.CountSyllablesPerWord(words)
	output := formula(ASL, ASW)
	return float32(math.Round(float64(output)*100) / 100)
}

func formula(averageSentenceLength float32, averageSyllablesPerWord float32) float32 {
	return 206.835 - (1.015 * averageSentenceLength) - (84.6 * averageSyllablesPerWord)
}
