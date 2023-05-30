package main

import (
	"fmt"
)

func main() {
	var salario float64
	var percentual float64

	// Leitura do salário mensal
	fmt.Print("Digite o salário mensal: ")
	fmt.Scanln(&salario)

	// Leitura do percentual de reajuste
	fmt.Print("Digite o percentual de reajuste: ")
	fmt.Scanln(&percentual)

	// Cálculo do novo salário
	novoSalario := salario + (salario * percentual / 100)

	// Exibição do novo salário
	fmt.Printf("Novo salário após o reajuste: R$ %.2f\n", novoSalario)
}
