package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Digite uma string:")
	var input string
	fmt.Scanln(&input)

	isNumeric := checkNumeric(input)

	if isNumeric {
		fmt.Println("A string contém somente números.")
	} else {
		fmt.Println("A string não contém somente números.")
	}
}

func checkNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
