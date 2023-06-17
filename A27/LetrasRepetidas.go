package main

import (
	"fmt"
	"strings"
)

func main() {
	var texto string
	fmt.Print("Digite o texto: ")
	fmt.Scanln(&texto)
	//teste

	texto = strings.ReplaceAll(texto, " ", "")

	frequencia := make(map[rune]int)

	for _, char := range texto {
		frequencia[char]++
	}

	for _, char := range texto {
		if frequencia[char] == 1 {
			fmt.Printf("A primeira letra não repetida é: %c\n", char)
			return
		}
	}

	fmt.Println("Não existem letras que não se repetem no texto informado.")
}
