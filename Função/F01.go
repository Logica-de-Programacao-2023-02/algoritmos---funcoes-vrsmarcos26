//Crie uma função que receba um slice de inteiros e retorne a média aritmética dos valores.

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	media := cm(slice)
	fmt.Println("Média dos valores:", media)
}

func cm(slice []int) float64 {
	if len(slice) == 0 {
		return 0.0
	}
	soma := 0
	for _, valor := range slice {
		soma += valor
	}

	return float64(soma) / float64(len(slice))
}
