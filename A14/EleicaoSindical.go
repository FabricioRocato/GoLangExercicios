package main

import "fmt"

func main() {
	var votosA, votosB, votosC, votosNulos, votosBranco int

	fmt.Print("Digite a quantidade de votos válidos para o candidato A: ")
	fmt.Scanln(&votosA)

	fmt.Print("Digite a quantidade de votos válidos para o candidato B: ")
	fmt.Scanln(&votosB)

	fmt.Print("Digite a quantidade de votos válidos para o candidato C: ")
	fmt.Scanln(&votosC)

	fmt.Print("Digite a quantidade de votos nulos: ")
	fmt.Scanln(&votosNulos)

	fmt.Print("Digite a quantidade de votos em branco: ")
	fmt.Scanln(&votosBranco)

	totalEleitores := votosA + votosB + votosC + votosNulos + votosBranco
	percentualValidos := float64((votosA + votosB + votosC)) / float64(totalEleitores) * 100
	percentualA := float64(votosA) / float64(totalEleitores) * 100
	percentualB := float64(votosB) / float64(totalEleitores) * 100
	percentualC := float64(votosC) / float64(totalEleitores) * 100
	percentualNulos := float64(votosNulos) / float64(totalEleitores) * 100
	percentualBranco := float64(votosBranco) / float64(totalEleitores) * 100

	fmt.Println("Resultado da Apuração:")
	fmt.Println("Número total de eleitores:", totalEleitores)
	fmt.Printf("Percentual de votos válidos em relação à quantidade de eleitores: %.2f%%\n", percentualValidos)
	fmt.Printf("Percentual de votos válidos do candidato A em relação à quantidade de eleitores: %.2f%%\n", percentualA)
	fmt.Printf("Percentual de votos válidos do candidato B em relação à quantidade de eleitores: %.2f%%\n", percentualB)
	fmt.Printf("Percentual de votos válidos do candidato C em relação à quantidade de eleitores: %.2f%%\n", percentualC)
	fmt.Printf("Percentual de votos nulos em relação à quantidade de eleitores: %.2f%%\n", percentualNulos)
	fmt.Printf("Percentual de votos em branco em relação à quantidade de eleitores: %.2f%%\n", percentualBranco)
}
