package main

import (
	"fmt"
	"strings"
)

func verificarJokenpo(jogador1, jogador2 string) string {
	jogador1 = strings.ToLower(jogador1)
	jogador2 = strings.ToLower(jogador2)

	vencedor := map[string]string{
		"tesourapapel":   "Jogador 1 venceu!",
		"papelpedra":     "Jogador 1 venceu!",
		"pedralagarto":   "Jogador 1 venceu!",
		"lagartospock":   "Jogador 1 venceu!",
		"spocktesoura":   "Jogador 1 venceu!",
		"tesouralagarto": "Jogador 1 venceu!",
		"lagartopapel":   "Jogador 1 venceu!",
		"papelspock":     "Jogador 1 venceu!",
		"spockpedra":     "Jogador 1 venceu!",
		"pedratesoura":   "Jogador 1 venceu!",
	}

	if jogador1 == jogador2 {
		return "Empate"
	}

	chave := jogador1 + jogador2
	if resultado, ok := vencedor[chave]; ok {
		return resultado
	}

	return "Jogador 2 venceu!"
}

func main() {
	var jogador1, jogador2 string
	fmt.Print("Jogador 1, escolha entre Pedra, Papel, Tesoura, Lagarto ou Spock: ")
	fmt.Scanln(&jogador1)
	fmt.Print("Jogador 2, escolha entre Pedra, Papel, Tesoura, Lagarto ou Spock: ")
	fmt.Scanln(&jogador2)

	resultado := verificarJokenpo(jogador1, jogador2)
	fmt.Println(resultado)
}
