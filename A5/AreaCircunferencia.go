package main

import "fmt"

func main() {
	var raio float64
	pi := 3.14159265

	fmt.Print("Digite o valor do raio da circunferência: ")
	fmt.Scanln(&raio)

	area := pi * raio * raio

	fmt.Println("A área da circunferência é:", area)
}
