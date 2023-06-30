package main

import (
	"fmt"
)

func main() {
	fmt.Println("Digite uma string:")
	var input string
	fmt.Scanln(&input)

	reversedString := reverseString(input)

	fmt.Println("String invertida:", reversedString)
}

func reverseString(s string) string {
	runes := []rune(s)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
