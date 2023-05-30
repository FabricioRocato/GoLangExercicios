package main

import (
	"fmt"
)

func main() {
	var numeros []float64

	for {
		var num float64
		fmt.Print("Digite um número (ou digite 0 para encerrar): ")
		fmt.Scanln(&num)

		if num == 0 {
			break
		}

		numeros = append(numeros, num)
	}

	total := 0.0
	for _, num := range numeros {
		total += num
	}

	media := total / float64(len(numeros))
	fmt.Printf("A média dos números é: %.2f\n", media)
}
