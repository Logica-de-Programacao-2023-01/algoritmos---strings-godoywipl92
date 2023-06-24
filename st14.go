package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Digite uma string:")
	var input string
	fmt.Scanln(&input)

	isDescending := checkDescending(input)

	if isDescending {
		fmt.Println("A string é uma sequência numérica decrescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica decrescente.")
	}
}

func checkDescending(s string) bool {
	length := len(s)
	if length <= 1 {
		return false
	}

	previousNum := 10
	for i := 0; i < length; i++ {
		num, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return false
		}

		if num >= previousNum {
			return false
		}

		previousNum = num
	}

	return true
}
