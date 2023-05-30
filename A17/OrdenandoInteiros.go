package main

import (
	"fmt"
	"sort"
)

func main() {
	var numeros [12]int

	// Leitura dos elementos
	for i := 0; i < len(numeros); i++ {
		fmt.Printf("Digite o elemento %d: ", i+1)
		fmt.Scanln(&numeros[i])
	}

	// Ordenação em ordem decrescente
	sort.Sort(sort.Reverse(sort.IntSlice(numeros[:])))

	// Apresentação dos elementos ordenados
	fmt.Println("Elementos em ordem decrescente:")
	for _, num := range numeros {
		fmt.Println(num)
	}
}
