package main

import (
	"fmt"
	"strings"
)

func isPangram(frase string) bool {
	frase = strings.ReplaceAll(frase, " ", "")
	frase = strings.ToLower(frase)

	letras := make(map[rune]bool)

	for _, char := range frase {
		if char >= 'a' && char <= 'z' {
			letras[char] = true
		}
	}

	for char := 'a'; char <= 'z'; char++ {
		if !letras[char] {
			return false
		}
	}

	return true
}

func main() {
	var frase string
	fmt.Print("Digite uma frase: ")
	fmt.Scanln(&frase)

	if isPangram(frase) {
		fmt.Println("A frase é um pangrama.")
	} else {
		fmt.Println("A frase não é um pangrama.")
	}
}
