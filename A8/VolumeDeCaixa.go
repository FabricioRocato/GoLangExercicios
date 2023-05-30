package main

import "fmt"

func main() {
	var comprimento, largura, altura float64

	fmt.Print("Digite o valor do comprimento da caixa: ")
	fmt.Scanln(&comprimento)

	fmt.Print("Digite o valor da largura da caixa: ")
	fmt.Scanln(&largura)

	fmt.Print("Digite o valor da altura da caixa: ")
	fmt.Scanln(&altura)

	volume := comprimento * largura * altura

	fmt.Printf("O volume da caixa é: %.2f\n", volume)
}
