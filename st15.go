package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Digite uma string:")
	var input string
	fmt.Scanln(&input)

	result := replaceVowels(input, "*")

	fmt.Println("Resultado:", result)
}

func replaceVowels(s, replacement string) string {
	vowels := "aeiouAEIOU"
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(vowels, r) {
			return []rune(replacement)[0]
		}
		return r
	}, s)
}
