package main

import "fmt"

func main() {
	var A, B int

	// Leitura dos valores para as variáveis A e B
	fmt.Print("Digite o valor de A: ")
	fmt.Scanln(&A)
	fmt.Print("Digite o valor de B: ")
	fmt.Scanln(&B)

	// Troca dos valores
	A, B = B, A

	// Apresentação dos valores após a troca
	fmt.Println("Valores após a troca:")
	fmt.Println("A =", A)
	fmt.Println("B =", B)
}
