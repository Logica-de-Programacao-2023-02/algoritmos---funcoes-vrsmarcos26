//Crie uma função que receba um número inteiro e um slice de inteiros como parâmetros e retorne verdadeiro
//se o número estiver presente no slice e falso caso contrário. Caso o slice esteja vazio, retorne um erro.

package main

import (
	"errors"
	"fmt"
)

func isNumberInSlice(number int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("O slice está vazio")
	}

	for _, n := range slice {
		if n == number {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	// Teste da função
	slice := []int{1, 2, 3, 4, 5}
	numberToFind := 3

	isPresent, err := isNumberInSlice(numberToFind, slice)

	if err != nil {
		fmt.Println("Erro:", err)
	} else if isPresent {
		fmt.Println(numberToFind, "está presente no slice.")
	} else {
		fmt.Println(numberToFind, "não está presente no slice.")
	}
}
