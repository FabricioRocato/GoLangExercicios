package main

import "fmt"

func main() {
	var inicio, fim int

	fmt.Print("Digite o número inicial do intervalo: ")
	fmt.Scanln(&inicio)

	fmt.Print("Digite o número final do intervalo: ")
	fmt.Scanln(&fim)

	somaPares := 0

	for num := inicio; num <= fim; num++ {
		if num%2 == 0 {
			somaPares += num
		}
	}

	fmt.Printf("A soma dos números pares no intervalo [%d, %d] é: %d\n", inicio, fim, somaPares)
}
