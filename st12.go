package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Digite uma string:")
	var input string
	fmt.Scanln(&input)

	isPalindrome := checkPalindrome(input)

	if isPalindrome {
		fmt.Println("A string é um palíndromo.")
	} else {
		fmt.Println("A string não é um palíndromo.")
	}
}

func checkPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = removeNonAlphanumeric(s)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func removeNonAlphanumeric(s string) string {
	var result strings.Builder

	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			result.WriteRune(r)
		}
	}

	return result.String()
}
