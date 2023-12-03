//Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice com os valores
//ordenados de forma crescente. Caso o slice esteja vazio, retorne um erro.

package main

import (
	"fmt"
)

func main() {
	slice := []int{5, 2, 9, 1, 7}
	novoSlice, err := ordenarCrescente(slice)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Valores ordenados em ordem crescente:", novoSlice)
	}
}

func ordenarCrescente(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, fmt.Errorf("O slice está vazio")
	}

	novoSlice := make([]int, len(slice))
	copy(novoSlice, slice)

	for i := 0; i < len(novoSlice); i++ {
		for j := i + 1; j < len(novoSlice); j++ {
			if novoSlice[i] > novoSlice[j] {
				novoSlice[i], novoSlice[j] = novoSlice[j], novoSlice[i]
			}
		}
	}

	return novoSlice, nil
}
