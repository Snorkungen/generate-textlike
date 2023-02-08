package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	text := generateSentence(25)

	fmt.Printf("%v \n%v \n", text, len(text))

}

func generateWord(len uint) string {
	if len < 2 {
		return Vowel.Get()
	}

	startsWithConsonant := len&1 == 1
	word := ""

	for i := 0; i < int(len); i++ {
		if startsWithConsonant && i&1 == 0 {
			word += Vowel.Get()
		} else {
			word += Consonant.Get()
		}
	}

	return strings.ToLower(word)
}

func capitalize(input string) string {
	return strings.ToUpper(input[:1]) + input[1:]
}

const SEED = "782814554715687"

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

	for len(sentence) < int(length) {
		max := int(length) - len(sentence) - 1
		wordLength := getWordLength(uint(max))

		sentence += fmt.Sprintf("%v ", generateWord(wordLength))
	}

	return capitalize(sentence)
}
