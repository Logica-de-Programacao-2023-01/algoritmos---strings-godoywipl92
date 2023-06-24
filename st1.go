package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Digite uma string:")
	var input string
	fmt.Scanln(&input)

	output := strings.ToUpper(input)

	fmt.Println("Resultado:", output)
}
