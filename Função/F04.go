//Crie uma função que receba um slice de inteiros e retorne o segundo maior valor.

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	r := m(slice)
	fmt.Print("Maior número da slice é: ", r)
}

func m(slice []int) int {
	maior := slice[0]
	r := slice[1]

	for _, n := range slice {
		if n > maior {
			r = maior
			maior = n
		} else if n > r && n != maior {
			r = n
		}
	}
	return r
}
