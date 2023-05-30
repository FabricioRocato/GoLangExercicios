package main

import (
	"fmt"
)

func isPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}

	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scanln(&numero)

	if isPrimo(numero) {
		fmt.Printf("%d é um número primo.\n", numero)
	} else {
		fmt.Printf("%d não é um número primo.\n", numero)
	}
}
