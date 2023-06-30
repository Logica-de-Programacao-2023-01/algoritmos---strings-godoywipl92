package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Digite uma string:")
	var input string
	fmt.Scanln(&input)

	words := strings.Fields(input)
	wordCount := len(words)

	fmt.Println("Número de palavras:", wordCount)
}
