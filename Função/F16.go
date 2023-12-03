//Escreva uma função que receba uma string como parâmetro e retorne uma nova string com
//todas as vogais substituídas por "*". Caso a string seja vazia, retorne um erro.

package main

import (
	"errors"
	"fmt"
	"strings"
)

func replaceVowelsWithAsterisk(input string) (string, error) {
	if input == "" {
		return "", errors.New("A string está vazia")
	}

	vowels := "aeiouAEIOU"
	for _, v := range vowels {
		input = strings.ReplaceAll(input, string(v), "*")
	}

	return input, nil
}

func main() {
	inputString := "Hello, World!"
	result, err := replaceVowelsWithAsterisk(inputString)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("String com vogais substituídas:", result)
	}
}
