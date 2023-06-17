package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var str1, str2 string
	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)
	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	// Remover espaços em branco das strings
	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")

	if len(str1) != len(str2) {
		fmt.Println("As strings não são anagramas.")
		return
	}

	slice1 := strings.Split(strings.ToLower(str1), "")
	slice2 := strings.Split(strings.ToLower(str2), "")

	sort.Strings(slice1)
	sort.Strings(slice2)

	if strings.Join(slice1, "") == strings.Join(slice2, "") {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}
