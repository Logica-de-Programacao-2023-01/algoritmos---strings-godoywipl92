package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println("Digite uma string:")
	var input string
	fmt.Scanln(&input)

	isCamelCase, wordCount := checkCamelCase(input)

	if isCamelCase {
		fmt.Println("A string está em camelCase.")
		fmt.Println("Número de palavras:", wordCount)
	} else {
		fmt.Println("A string não está em camelCase.")
	}
}

func checkCamelCase(s string) (bool, int) {
	if s == "" {
		return false, 0
	}

	wordCount := 1

	for i, char := range s {
		if i == 0 && !unicode.IsUpper(char) {
			return false, 0
		} else if i != 0 && unicode.IsUpper(char) {
			wordCount++
		}
	}

	return true, wordCount
}
