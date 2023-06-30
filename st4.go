package main

import "fmt"

func main() {
	fmt.Println("Digite a primeira string:")
	var str1 string
	fmt.Scanln(&str1)

	fmt.Println("Digite a segunda string:")
	var str2 string
	fmt.Scanln(&str2)

	if str1 == str2 {
		fmt.Println("As strings são iguais.")
	} else {
		fmt.Println("As strings são diferentes.")
	}
}
