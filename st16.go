package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Digite a primeira string:")
	var str1 string
	fmt.Scanln(&str1)

	fmt.Println("Digite a segunda string:")
	var str2 string
	fmt.Scanln(&str2)

	isSubstring := checkSubstring(str1, str2)

	if isSubstring {
		fmt.Println("A segunda string é uma substring da primeira.")
	} else {
		fmt.Println("A segunda string não é uma substring da primeira.")
	}
}

func checkSubstring(str1, str2 string) bool {
	return strings.Contains(str1, str2)
}
