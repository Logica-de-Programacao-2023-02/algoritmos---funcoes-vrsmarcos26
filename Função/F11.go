//Crie uma função que receba um número inteiro como parâmetro e retorne verdadeiro se ele for um número
//primo e falso caso contrário. Caso o número seja negativo, retorne um erro.

package main

import (
	"errors"
	"fmt"
	"math"
)

func isPrime(n int) (bool, error) {
	if n < 0 {
		return false, errors.New("Número negativo não é permitido")
	}

	if n <= 1 {
		return false, nil
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	num := 17
	prime, err := isPrime(num)

	if err != nil {
		fmt.Println("Erro:", err)
	} else if prime {
		fmt.Println(num, "é um número primo.")
	} else {
		fmt.Println(num, "não é um número primo.")
	}
}
