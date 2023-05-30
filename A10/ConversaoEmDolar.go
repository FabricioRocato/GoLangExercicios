package main

import "fmt"

func main() {
	var cotacao, reais float64

	fmt.Print("Digite a cotação do dólar: ")
	fmt.Scanln(&cotacao)

	fmt.Print("Digite a quantidade de reais: ")
	fmt.Scanln(&reais)

	conversao := reais / cotacao

	fmt.Printf("O valor em dólares é: US$ %.2f\n", conversao)
}
