package main

import "fmt"

func main() {
	var valorHora, horasTrabalhadas, descontoINSS float64

	fmt.Print("Digite o valor da hora-aula: ")
	fmt.Scanln(&valorHora)

	fmt.Print("Digite o número de horas trabalhadas no mês: ")
	fmt.Scanln(&horasTrabalhadas)

	fmt.Print("Digite o percentual de desconto do INSS: ")
	fmt.Scanln(&descontoINSS)

	salarioBruto := valorHora * horasTrabalhadas
	desconto := salarioBruto * (descontoINSS / 100)
	salarioLiquido := salarioBruto - desconto

	fmt.Printf("Salário bruto: R$ %.2f\n", salarioBruto)
	fmt.Printf("Salário líquido: R$ %.2f\n", salarioLiquido)
}
