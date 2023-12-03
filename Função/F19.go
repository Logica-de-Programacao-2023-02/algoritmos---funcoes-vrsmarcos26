//Crie uma função que receba um número inteiro como parâmetro e retorne um novo slice contendo
//todos os números primos menores ou iguais a ele. Caso o número seja negativo, retorne um erro.

package main

import (
	"errors"
	"fmt"
	"math"
)

func generatePrimesUpToN(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("Número negativo não é permitido")
	}

	var primes []int

	for i := 2; i <= n; i++ {
		if primo(i) {
			primes = append(primes, i)
		}
	}

	return primes, nil
}

func primo(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	num := 20
	primes, err := generatePrimesUpToN(num)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números primos até", num, "são:", primes)
	}
}
