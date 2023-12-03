//Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice
//com apenas os números pares contidos no slice. Caso o slice esteja vazio, retorne um erro.

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	nvslice, err := afuncao(slice, par)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Novo slice:", nvslice)
	}
}

func afuncao(slice []int, funcao func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, fmt.Errorf("O slice está vazio")
	}

	novoSlice := make([]int, 0)
	for _, valor := range slice {
		resultado := funcao(valor)
		if resultado != 0 {
			novoSlice = append(novoSlice, resultado)
		}
	}
	return novoSlice, nil
}

func par(x int) int {
	if x%2 == 0 {
		return x
	}
	return 0
}
