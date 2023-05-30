package main

import (
	"fmt"
	"strings"
)

func converterRomanosParaInteiro(numeralRomano string) int {
	valores := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}

	total := 0
	numeralRomano = strings.ToUpper(numeralRomano)

	for i := 0; i < len(numeralRomano); i++ {
		valorAtual := valores[string(numeralRomano[i])]
		if i+1 < len(numeralRomano) {
			valorSeguinte := valores[string(numeralRomano[i+1])]
			if valorAtual >= valorSeguinte {
				total += valorAtual
			} else {
				total -= valorAtual
			}
		} else {
			total += valorAtual
		}
	}

	return total
}

func main() {
	var numeralRomano string
	fmt.Print("Digite um numeral romano: ")
	fmt.Scanln(&numeralRomano)

	numeroInteiro := converterRomanosParaInteiro(numeralRomano)
	fmt.Printf("O numeral romano %s é equivalente a %d\n", numeralRomano, numeroInteiro)
}
