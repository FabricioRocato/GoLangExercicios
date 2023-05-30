package main

import "fmt"

func main() {
	var cotacao, dolares float64

	fmt.Print("Digite a cotação do dólar: ")
	fmt.Scanln(&cotacao)

	fmt.Print("Digite a quantidade de dólares: ")
	fmt.Scanln(&dolares)

	conversao := cotacao * dolares

	fmt.Printf("O valor em reais é: R$ %.2f\n", conversao)
}
