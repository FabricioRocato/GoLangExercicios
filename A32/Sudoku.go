package main

import (
	"fmt"
)

func validarSudoku(tabuleiro [][]int) (bool, [][]bool) {
	for i := 0; i < 9; i++ {
		linhaValida := make([]bool, 9)
		colunaValida := make([]bool, 9)

		for j := 0; j < 9; j++ {
			linhaValor := tabuleiro[i][j]
			colunaValor := tabuleiro[j][i]

			if linhaValor != 0 {
				if linhaValida[linhaValor-1] {
					return false, nil
				}
				linhaValida[linhaValor-1] = true
			}

			if colunaValor != 0 {
				if colunaValida[colunaValor-1] {
					return false, nil
				}
				colunaValida[colunaValor-1] = true
			}
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			regiaoValida := make([]bool, 9)

			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					regiaoValor := tabuleiro[k][l]

					if regiaoValor != 0 {
						if regiaoValida[regiaoValor-1] {
							return false, nil
						}
						regiaoValida[regiaoValor-1] = true
					}
				}
			}
		}
	}

	return true, nil
}

//teste

func main() {
	tabuleiro := make([][]int, 9)
	for i := 0; i < 9; i++ {
		tabuleiro[i] = make([]int, 9)
	}

	tabuleiro[0][0] = 5
	tabuleiro[0][1] = 3
	tabuleiro[0][4] = 6
	tabuleiro[1][0] = 6
	tabuleiro[1][3] = 1
	tabuleiro[1][4] = 9
	tabuleiro[1][5] = 5
	tabuleiro[2][1] = 9
	tabuleiro[2][2] = 8
	tabuleiro[2][7] = 6
	tabuleiro[3][0] = 8
	tabuleiro[3][4] = 6
	tabuleiro[3][8] = 3
	tabuleiro[4][0] = 4
	tabuleiro[4][3] = 8
	tabuleiro[4][5] = 3
	tabuleiro[4][8] = 1
	tabuleiro[5][0] = 7
	tabuleiro[5][4] = 2
	tabuleiro[5][8] = 6
	tabuleiro[6][1] = 6
	tabuleiro[6][6] = 2
	tabuleiro[6][7] = 8
	tabuleiro[7][3] = 4
	tabuleiro[7][4] = 1
	tabuleiro[7][5] = 9
	tabuleiro[7][8] = 5
	tabuleiro[8][4] = 8
	tabuleiro[8][7] = 7
	tabuleiro[8][8] = 9

	valido, _ := validarSudoku(tabuleiro)

	if valido {
		fmt.Println("O tabuleiro Sudoku está válido.")
	} else {
		fmt.Println("O tabuleiro Sudoku contém células com valores incorretos.")
	}
}
