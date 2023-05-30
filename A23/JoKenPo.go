package main

import (
	"fmt"
	"strings"
)

func verificarJokenpo(jogador1, jogador2 string) string {
	jogador1 = strings.ToLower(jogador1)
	jogador2 = strings.ToLower(jogador2)

	if jogador1 == jogador2 {
		return "Empate"
	}

	if (jogador1 == "pedra" && jogador2 == "tesoura") ||
		(jogador1 == "tesoura" && jogador2 == "papel") ||
		(jogador1 == "papel" && jogador2 == "pedra") {
		return "Jogador 1 venceu!"
	}

	return "Jogador 2 venceu!"
}

func main() {
	var jogador1, jogador2 string
	fmt.Print("Jogador 1, escolha Pedra, Papel ou Tesoura: ")
	fmt.Scanln(&jogador1)
	fmt.Print("Jogador 2, escolha Pedra, Papel ou Tesoura: ")
	fmt.Scanln(&jogador2)

	resultado := verificarJokenpo(jogador1, jogador2)
	fmt.Println(resultado)
}
