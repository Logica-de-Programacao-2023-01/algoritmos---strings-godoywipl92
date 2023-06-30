package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Digite a primeira string:")
	var str1 string
	fmt.Scanln(&str1)

	fmt.Println("Digite a segunda string:")
	var str2 string
	fmt.Scanln(&str2)

	isAnagram := checkAnagram(str1, str2)

	if isAnagram {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}

func checkAnagram(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	str1 = sortString(str1)
	str2 = sortString(str2)

	return str1 == str2
}

func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
