package fleschreading

// The Flesch score uses the number of syllables and sentence lengths to
// determine the reading ease of the sample.

// A Flesch score of 60 is taken to be plain English. A score in the range of 60-70 corresponds
// to 8th/9th grade English level. A score between 50 and 60 corresponds to a 10th/12th grade level.
// Below 30 is college graduate level. To give you a feel for what the different levels are like,
// most states require scores from 40 to 50 for insurance documents.

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

func formula(averageSentenceLength float32, AvgSyllablesPerWord float32) float32 {
	return 206.835 - (1.015 * averageSentenceLength) - (84.6 * AvgSyllablesPerWord)
}
