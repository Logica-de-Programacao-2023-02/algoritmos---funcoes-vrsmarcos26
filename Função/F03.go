//Crie uma função que receba um slice de strings e retorne a concatenação de todas as strings.

package main

import "fmt"

func main() {
	slice := []string{"a", "b", "c", "d"}
	r := cts(slice)
	fmt.Print("Seu slice concatenado fica ", r)
}

func cts(slice []string) string {
	soma := ""
	for _, l := range slice {
		soma += l
	}
	return soma
}
