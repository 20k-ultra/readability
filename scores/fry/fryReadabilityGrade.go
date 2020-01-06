package fry

// Fry Readability graph was developed by Edward Fry and is often
// selected for it's simplicity and accuracy. The graph has two axes: the average
// number of syllables (x-axis) and the average number of sentences (y-axis) per hundred words.
// Passages of text that are at least one hundreds words can be plotted on the
// graph to find the corresponding grade level.

import (
	"math"
	// "fmt"

	"github.com/20k-ultra/readability"
)

// Calculate implements Fry Readability formula
func Calculate(text []byte) float32 {
	words := readability.GetWords(text, false)
	if (len(words) < 300){
		panic("Fry Graph Readability requires atleast 300 words.")
	}
	output := formula(words)
	return float32(math.Round(float64(output)*100) / 100)
}

func formula(words []string,) int {
	samples := readability.GetWordSample(words, 100)
	sentenceCount := 0
	syllablesCount := 0
	for i := 0; i < 3; i++ { // get 3 random samples
		var index int
		var sample string
		if (len(samples) == 1){
			index = 0
			sample = samples[0]
		} else {
			index, sample = readability.GetRandomSample(samples)
		}
		samples = append(samples[:index], samples[index+1:]...)
		sentenceCount += readability.CountSentences(sample)
		syllablesCount += readability.CountSyllables(sample)
	}
	sentenceAverage := float64(sentenceCount) / 3
	syllablesAverage := float64(syllablesCount) / 3
	return getGrade(float32(math.Round(syllablesAverage)), float32(math.Round(sentenceAverage)))
}