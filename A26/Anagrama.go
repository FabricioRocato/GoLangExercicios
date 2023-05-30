package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Receber as duas strings informadas pelo usuário
	var str1, str2 string
	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)
	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	// Remover espaços em branco das strings
	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")

	// Verificar se as strings possuem o mesmo tamanho
	if len(str1) != len(str2) {
		fmt.Println("As strings não são anagramas.")
		return
	}

	// Converter as strings em slices de caracteres
	slice1 := strings.Split(strings.ToLower(str1), "")
	slice2 := strings.Split(strings.ToLower(str2), "")

	// Ordenar as slices
	sort.Strings(slice1)
	sort.Strings(slice2)

	// Verificar se as slices ordenadas são iguais
	if strings.Join(slice1, "") == strings.Join(slice2, "") {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}
