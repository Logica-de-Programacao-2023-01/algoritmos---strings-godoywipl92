package main

import (
	"fmt"
)

func main() {
	fmt.Println("Digite uma string:")
	var input string
	fmt.Scanln(&input)

	uniqueLetters := getUniqueLetters(input)

	fmt.Println("Letras únicas:", uniqueLetters)
}

func getUniqueLetters(s string) string {
	letterCount := make(map[rune]int)

 for _, letter := range s {
		letterCount[letter]++
	}

	uniqueLetters := ""

	for _, letter := range s {
		if letterCount[letter] == 1 {
			uniqueLetters += string(letter)
		}
	}

	return uniqueLetters
}
