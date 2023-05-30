package main

import "fmt"

func calcularFatorial(numero int) int {
	if numero == 0 {
		return 1
	}

	fatorial := 1
	for i := 1; i <= numero; i++ {
		fatorial *= i
	}

	return fatorial
}

func main() {
	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scanln(&numero)

	fatorial := calcularFatorial(numero)
	fmt.Printf("O fatorial de %d é %d\n", numero, fatorial)
}
