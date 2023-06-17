package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func breakLines(frase string, colunas int) []string {
	palavras := strings.Split(frase, " ")
	linhas := []string{}
	linhaAtual := ""

	for _, palavra := range palavras {
		if len(linhaAtual)+len(palavra)+1 <= colunas {
			linhaAtual += palavra + " "
		} else {
			linhas = append(linhas, strings.TrimSpace(linhaAtual))
			linhaAtual = palavra + " "
		}
	}

	linhas = append(linhas, strings.TrimSpace(linhaAtual))
	return linhas
}

//teste

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite a frase: ")
	frase, _ := reader.ReadString('\n')

	fmt.Print("Digite a quantidade de colunas: ")
	var colunas int
	fmt.Scanln(&colunas)

	linhasQuebradas := breakLines(frase, colunas)

	for _, linha := range linhasQuebradas {
		fmt.Println(linha)
	}
}
