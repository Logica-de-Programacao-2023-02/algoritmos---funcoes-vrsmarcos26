//Crie uma função que receba uma string e retorne a quantidade de vogais presentes.

package main

import (
	"fmt"
)

func main() {
	str := "Exemplo de frase com vogais"
	r := n(str)
	fmt.Print("Quantidade de vogal é igual a ", r)
}

func n(str string) int {
	soma := 0
	for _, vogal := range str {
		if vogal == 'a' || vogal == 'e' || vogal == 'i' || vogal == 'o' || vogal == 'u' || vogal == 'A' || vogal == 'E' || vogal == 'I' || vogal == 'O' || vogal == 'U' {
			soma++
		}
	}
	return soma
}
