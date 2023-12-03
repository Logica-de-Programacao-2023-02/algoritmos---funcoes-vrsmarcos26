//Escreva uma função que receba um slice de strings como parâmetro e retorne uma string contendo
//apenas as strings que começam com uma letra maiúscula. Caso o slice esteja vazio, retorne um erro.

package main

import (
	"errors"
	"fmt"
	"unicode"
)

func filterUpperCase(strings []string) (string, error) {
	if len(strings) == 0 {
		return "", errors.New("Slice está vazio")
	}

	var result string

	for _, s := range strings {
		if len(s) > 0 && unicode.IsUpper(rune(s[0])) {
			result += s + " "
		}
	}

	return result[:len(result)-1], nil
}

func main() {
	stringSlice := []string{"Apple", "banana", "Orange", "grape"}
	filteredString, err := filterUpperCase(stringSlice)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings com letra maiúscula:", filteredString)
	}
}
