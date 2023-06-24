package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Digite uma string:")
	var input string
	fmt.Scanln(&input)

	result := removeVowels(input)

	fmt.Println("Resultado:", result)
}

func removeVowels(s string) string {
	vowels := "aeiouAEIOU"
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(vowels, r) {
			return -1
		}
		return r
	}, s)
}
