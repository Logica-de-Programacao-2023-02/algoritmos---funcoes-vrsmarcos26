//Escreva uma função que receba um slice de inteiros e uma função como parâmetros.
//A função deve aplicar a função recebida em cada elemento do slice e retornar a soma de todos os resultados.
//Caso o slice esteja vazio, retorne um erro.

package main

import (
	"errors"
	"fmt"
)

func applyFunctionAndSum(slice []int, applyFunc func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("O slice está vazio")
	}

	sum := 0
	for _, num := range slice {
		result := applyFunc(num)
		sum += result
	}

	return sum, nil
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	squareFunc := func(x int) int {
		return x * x
	}

	result, err := applyFunctionAndSum(intSlice, squareFunc)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Soma dos resultados da função aplicada:", result)
	}
}
