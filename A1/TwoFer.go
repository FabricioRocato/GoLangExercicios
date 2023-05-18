package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TwoFer(name string) string {
	if name == "" {
		name = "você"
	}
	return fmt.Sprintf("Um para %s, um para mim!", name)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Informe um nome: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	result := TwoFer(name)
	fmt.Println(result)
}
