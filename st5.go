package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Digite um número em ponto flutuante:")
	var input string
	fmt.Scanln(&input)

	_, err := strconv.ParseFloat(input, 64)

	if err == nil {
		fmt.Println("É um número válido em ponto flutuante.")
	} else {
		fmt.Println("Não é um número válido em ponto flutuante.")
	}
}
