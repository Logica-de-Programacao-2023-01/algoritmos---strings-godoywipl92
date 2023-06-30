package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Digite uma string:")
	var input string
	fmt.Scanln(&input)

	fmt.Println("Digite a letra a ser substituída:")
	var oldLetter string
	fmt.Scanln(&oldLetter)

	fmt.Println("Digite a nova letra:")
	var newLetter string
	fmt.Scanln(&newLetter)

	replaced := strings.ReplaceAll(input, oldLetter, newLetter)

	fmt.Println("Resultado:", replaced)
}
