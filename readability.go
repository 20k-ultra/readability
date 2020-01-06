package readability

import (
	"strings"
	"regexp"
	"unicode/utf8"
	"math"
	"math/rand"
	"time"
	// "fmt"
)

type offset struct {
	start int
	end int
}

// GetSentences returns slice of strings representing
// sentences from a larger slice of bytes.
func GetSentences(text []byte) []string {
	sentences := strings.Split(string(text), ".")
	if (sentences[len(sentences) - 1] == "") {
		sentences = sentences[:len(sentences)-1] //pop last element since it's empty
	}
	return sentences
}

// GetWords returns slice of strings representing
// words from a larger slice of bytes.
func GetWords(text []byte, removePeriods bool) []string {
	wordsPattern := `[^\w\s]`
	if (!removePeriods){
		wordsPattern = `[^\w\s.]` // allow periods
	}
	regex := regexp.MustCompile(wordsPattern)
	normalized := regex.ReplaceAllString(string(text), "") //remove punctuation
	regex2 := regexp.MustCompile(`[\n]`)
	normalized = regex2.ReplaceAllString(normalized, " ") //remove new lines
	words := strings.Split(normalized, " ")
	if (words[len(words) - 1] == "") {
		words = words[:len(words)-1] //pop last element since it's empty
	}
	return words
}

// GetWordSample attempts to return a slice of strings which matches
// the number of words desired in the sample by sampleSize.
func GetWordSample(words []string, sampleSize int) []string {
	offsetIndexes := calcOffsetIndexes(len(words), sampleSize)
	samples := make([]string, 0)
	for _, offset := range offsetIndexes {
		samples = append(samples, GetWordsInRange(words, offset.start, offset.end))
	}
	return samples
}

// GetWordsWithSyllables returns slice of strings that contains the
// required amount of syllables
func GetWordsWithSyllables(words []string, amountOfSyllables int) []string {
	hasEnoughSyllables := make([]string, 0)
	for _, word := range words {
		if (CountSyllables(word) >= amountOfSyllables){
			hasEnoughSyllables = append(hasEnoughSyllables, word)
		}
	}
	return hasEnoughSyllables
}

// GetWordsInRange returns a concatted string of words from slice of strings within range (start to end)
func GetWordsInRange(words []string, start int, end int) string {
	wordSlice := make([]string, 0)
	for i := start; i <= end; i++ {
		wordSlice = append(wordSlice, words[i])
	}
	return strings.Join(wordSlice, " ")
}

// CountWordsPerSentence returns float32 average of words per sentence.
func CountWordsPerSentence(sentences []string) float32 {
	sentenceCount := len(sentences)
	entireText := strings.Join(sentences, " ")
	words := GetWords([]byte(entireText), true)
	fraction := float64(len(words)) / float64(sentenceCount)
	return float32(math.Round(fraction*100)/100)
}

// CountSentences returns int sum of sentences found in a piece of text.
func CountSentences(text string) int {
	sentences := strings.Split(text, ".")
	if (sentences[len(sentences) - 1] == "") {
		sentences = sentences[:len(sentences)-1] //pop last element since it's empty
	}
	return len(sentences)
}

// CountCharactersPerWord returns float32 average of characters per word.
func CountCharactersPerWord(words []string) float32 {
	charCount := utf8.RuneCountInString(strings.Join(words, ""))
	fraction := float64(charCount) / float64(len(words))
	return float32(math.Round(fraction*100)/100)
}

// CountSyllablesInSlice returns int value of syllables counted in slice of strings
func CountSyllablesInSlice(words []string) int {
	syllablesCount := 0
	for _, word := range words {
		syllablesCount += CountSyllables(word)
	}
	return syllablesCount
}

// CountSyllablesPerWord returns float32 average of syllables per word.
func CountSyllablesPerWord(words []string) float32 {
	wordCount := len(words)
	syllablesCount := CountSyllablesInSlice(words)
	return float32(syllablesCount) / float32(wordCount)
}

// CountSyllables returns int representing the number of syllables in a string.
func CountSyllables(word string) int {
	word = strings.ToLower(word)
	regex1 := regexp.MustCompile(`(?:[^laeiouy]|ed|[^laeiouy]e)$`)
	firstNormalize := regex1.ReplaceAllString(word, "")
	regex2 := regexp.MustCompile(`^y`)
	secondNormalize := regex2.ReplaceAllString(firstNormalize, "")
	regex3 := regexp.MustCompile(`[aeiouy]{1,2}`)
	syl := regex3.FindAllString(secondNormalize, -1)
	if (len(syl) == 0){
		return 1
	}
	return len(syl)
}

// IsProperNoun returns boolean if the word is a proper noun
// WARNING: This is a niave approach to checking if it's a proper noun.
// See: https://en.wikipedia.org/wiki/Part-of-speech_tagging
// We are just checking if the first character is upper case
func IsProperNoun(word string) bool {
	regex := regexp.MustCompile(`^[A-Z]`)
	upperCase := regex.FindAllString(word, -1)
	return len(upperCase) > 0
}

// AverageSentenceLength returns float32 value of the ASL (average sentence length).
func AverageSentenceLength(words []string, sentences []string) float32 {
	return float32(len(words)) / float32(len(sentences))
}

// calcOffsetIndexes returns slice of offsets that have the start/end
// indexes required to evenly sample a slice.
func calcOffsetIndexes(size int, length int) []offset {
	divides := float64(size) / float64(length)
	roundedUpDivides := int(math.Floor(divides))
	offsets := make([]offset, 0)
	position := 0
	end := length - 1
	for i := 0; i < roundedUpDivides; i++ {
		offsets = append(offsets, offset{
			start: position,
			end:   end,
		})
		position = end + 1
		end += length
	}
	return offsets
}

// GetRandomSample returns a randomly selected item from slice of strings
func GetRandomSample(items []string) (int, string) {
	if len(items) == 1 {
		panic("Getting random item from selection of 1 item cannot be random.")
	}
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	position := r.Intn(len(items))
	return position, items[position]
}