//Crie uma função que receba um slice de inteiros e um valor inteiro, e retorne a posição do
//primeiro elemento igual ao valor no slice. Caso não encontre, retorne -1.

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	valor := 3
	posicao := encontrarPosicao(slice, valor)

	if posicao == -1 {
		fmt.Printf("O valor %d não foi encontrado no slice.\n", valor)
	} else {
		fmt.Printf("O valor %d foi encontrado na posição %d do slice.\n", valor, posicao)
	}
}

func encontrarPosicao(slice []int, valor int) int {
	for i, elemento := range slice {
		if elemento == valor {
			return i
		}
	}
	return -1
}
