package main

import "fmt"

func fibonacci(n int) []int {
	sequence := make([]int, 0)

	if n <= 0 {
		return sequence
	}

	a, b := 0, 1
	sequence = append(sequence, a)

	for b <= n {
		sequence = append(sequence, b)
		a, b = b, a+b
	}

	return sequence
}

// teste
func main() {
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scanln(&num)

	result := fibonacci(num)

	fmt.Println("Sequência de Fibonacci até", num, ":")
	fmt.Println(result)
}
