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

	soma := num1 + num2 + num3
	quadradoSoma := soma * soma

	fmt.Println("O quadrado da soma dos valores é:", quadradoSoma)
}
