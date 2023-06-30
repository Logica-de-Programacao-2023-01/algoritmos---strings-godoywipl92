package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Digite uma string:")
	var input string
	fmt.Scanln(&input)

	fmt.Println("Digite o caractere a ser substituído:")
	var oldChar string
	fmt.Scanln(&oldChar)

	fmt.Println("Digite o novo caractere:")
	var newChar string
	fmt.Scanln(&newChar)

	output := strings.ReplaceAll(input, oldChar, newChar)

	fmt.Println("Resultado:", output)
}
