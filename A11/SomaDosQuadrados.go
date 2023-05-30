package main

import "fmt"

func main() {
	var num1, num2, num3 int

	fmt.Print("Digite o primeiro valor: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo valor: ")
	fmt.Scanln(&num2)

	fmt.Print("Digite o terceiro valor: ")
	fmt.Scanln(&num3)

	somaQuadrados := (num1 * num1) + (num2 * num2) + (num3 * num3)

	fmt.Println("A soma dos quadrados dos valores é:", somaQuadrados)
}
