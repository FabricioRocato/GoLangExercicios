package main

import "fmt"

func main() {
	var nota1, nota2, nota3, nota4 float64

	fmt.Print("Digite a primeira nota: ")
	fmt.Scanln(&nota1)

	fmt.Print("Digite a segunda nota: ")
	fmt.Scanln(&nota2)

	fmt.Print("Digite a terceira nota: ")
	fmt.Scanln(&nota3)

	fmt.Print("Digite a quarta nota: ")
	fmt.Scanln(&nota4)

	media := (nota1 + nota2 + nota3 + nota4) / 4

	fmt.Printf("Média do aluno: %.2f\n", media)

	if media >= 5 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}
