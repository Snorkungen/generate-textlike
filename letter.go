package main

type Letter struct {
	index   uint
	options string
}

func (letter *Letter) Get() string {
	char := string(letter.options[letter.index])
	letter.index = (letter.index + 1) % uint(len(letter.options))
	return char
}

var Vowel = Letter{
	options: "AEOUEI",
}

var Consonant = Letter{
	options: "WQSRTDZXCVBFGTYHNMJKL",
}
