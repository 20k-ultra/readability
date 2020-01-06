package gunningfogscale

// The Gunning Fog scale is similar to the Flesch scale in that it uses syllable
// counts and sentence length. The scale uses the percentage of 'Foggy' words,
// those that contain 3 or more syllables.

// A fog score of 5 is readable, 10 is hard, 15 is difficult, and 20 is very difficult.

import (
	"math"
	"regexp"
	"fmt"

	"github.com/20k-ultra/readability"
)

// Calculate implements Gunning Fog Scale formula
func Calculate(text []byte) float32 {
	words := readability.GetWords(text, true)
	samples := readability.GetWordSample(words, 100)
	_, sample := readability.GetRandomSample(samples)
	sampleWords := readability.GetWords([]byte(sample), false)
	sampleSentences := readability.GetSentences([]byte(sample))
	twoSyllableWords := readability.GetWordsWithSyllables(sampleWords, 2)
	threeSyllableWords := readability.GetWordsWithSyllables(sampleWords, 3)
	hardWordCount := 0
	for _, word := range threeSyllableWords {
		if (!readability.IsProperNoun(word) && !hyphenatedWord(word)){
			hardWordCount++
		}
	}
	fmt.Printf("only threeSyllableWords: %d \n", hardWordCount)

	for _, word := range twoSyllableWords {
		if (!has2WordSuffix(word)){
			hardWordCount++
		}
	}

	fmt.Printf("and twoSyllableWords: %d \n", hardWordCount)

	ASL := readability.AverageSentenceLength(sampleWords, sampleSentences)
	PHW := float32(hardWordCount) / float32(len(words))
	output := formula(ASL, PHW)
	return float32(math.Round(float64(output)*100) / 100)
}

func formula(averageSentenceLength float32, percentHardWords float32) float32 {
	return 0.4 * (averageSentenceLength + percentHardWords)
}

func has2WordSuffix(word string) bool {
	regex := regexp.MustCompile(`(es|ed)\b`)
	suffix := regex.FindAllString(word, -1)
	return len(suffix) > 0
}

func hyphenatedWord(word string) bool {
	regex := regexp.MustCompile(`-`)
	hyphenes := regex.FindAllString(word, -1)
	return len(hyphenes) > 0
}