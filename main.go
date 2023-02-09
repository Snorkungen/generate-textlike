package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	text := generateSentence(65)

	fmt.Printf("%v \n%v \n", text, len(text))

}

var patterns = [4]string{"cvcvvcv", "vccvc", "cvvcvc", "cvvcvc"}
var patternIndex = 0

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
	return strings.ToUpper(input[:1]) + input[1:]
}

const SEED = "4673578782814554715687"

var seedIndex uint8 = 0

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
