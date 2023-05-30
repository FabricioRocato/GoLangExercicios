package main

import (
	"fmt"
)

func fizzBuzz(inicio, fim int) {
	if fim <= inicio {
		fmt.Println("Número final deve ser maior que o número inicial.")
		return
	}

	for i := inicio; i <= fim; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	var inicio, fim int

	fmt.Print("Digite o número inicial do intervalo: ")
	fmt.Scanln(&inicio)

	fmt.Print("Digite o número final do intervalo: ")
	fmt.Scanln(&fim)

	fizzBuzz(inicio, fim)
}
