package main

import (
	fmt "fmt"
	"strings"
)

func main() {
	var x string
	fmt.Print("Digite uma palavra, frase ou número: ")
	fmt.Scanln(&x)

	isPalindrome := checkPalindrome(x)
	fmt.Println("Palavra/Frase/Número:", x)
	fmt.Println("É palíndromo?", isPalindrome)
}

func checkPalindrome(x string) bool {
	// Remover espaços em branco e converter para letras minúsculas
	x = strings.ToLower(strings.ReplaceAll(x, " ", ""))

	// Percorrer a string de ambos os lados
	for i := 0; i < len(x)/2; i++ {
		if x[i] != x[len(x)-i-1] {
			return false
		}
	}
	return true
}
