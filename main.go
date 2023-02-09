package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_TEXT_LENGTH = 10000
const DEFAULT_SENTENCE_LENGTH uint = 48
const SEED = "4673578782814554715687"

var seedIndex uint8 = 0

// v = vowel c = consonant
var patterns = [4]string{"cvcvvcv", "vccvc", "cvvcvc", "cvvcvc"}
var patternIndex = 0

func main() {
	if len(os.Args) <= 1 {
		panic("Mising: please enter text of length.")
	}

	length, err := strconv.ParseUint(os.Args[1], 10, 64)

	if err != nil {
		panic("Failed: input length invalid.")
	}

	for length != 0 {
		text := generateText(uint(length % uint64(MAX_TEXT_LENGTH)))
		fmt.Print(text)

		length -= uint64(len(text))
	}
}

func getPattern() string {
	pattern := patterns[patternIndex]
	patternIndex = (patternIndex + 1) % len(patterns)
	return pattern
}

func generateWord(length uint) string {
	if length < 2 {
		return strings.ToLower(Vowel.Get())
	}

	pattern := getPattern()
	word := ""

	for i := 0; i < int(length); i++ {
		v := string(pattern[i%len(pattern)])
		if v == "c" {
			word += Consonant.Get()
		} else {
			word += Vowel.Get()
		}
	}

	return strings.ToLower(word)
}

func capitalize(input string) string {
	if len(input) < 1 {
		return input
	}
	return strings.ToUpper(input[:1]) + input[1:]
}

func getWordLength(max uint) uint {
	length, err := strconv.ParseUint(string(SEED[seedIndex]), 10, 4)

	if err != nil || uint(length) > max {
		return max
	}

	seedIndex = (seedIndex + 1) % uint8(len(SEED))

	return uint(length)
}

func generateSentence(length uint) string {
	if length <= 1 {
		return capitalize(generateWord((1)))
	}

	sentence := ""

	var wordLength uint
	var diff int

	for len(sentence) < int(length) {
		diff = int(length) - len(sentence)
		wordLength = getWordLength(uint(diff) - 1)

		if diff < int(wordLength)+1 {
			wordLength = uint(diff) - 1
		}
		if diff-(int(wordLength)+1) < 2 {
			wordLength = uint(diff) - 1
		}

		sentence += fmt.Sprintf("%v ", generateWord(wordLength))
	}

	sentence = sentence[:len(sentence)-1] + "."

	return capitalize(sentence)
}

func generateText(length uint) string {
	text := ""

	for {
		if len(text) == int(length) {
			break
		}

		sentenceLength := DEFAULT_SENTENCE_LENGTH

		if len(text) == 0 {
			sentenceLength += length % DEFAULT_SENTENCE_LENGTH
		} else {
			text += " "
		}

		diff := length - uint(len(text))

		if sentenceLength > diff {
			sentenceLength = diff
		}

		text += generateSentence(sentenceLength)
	}

	return text
}
