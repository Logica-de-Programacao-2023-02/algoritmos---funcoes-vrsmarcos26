//Crie uma função que receba uma string como parâmetro e retorne um novo slice com todas as palavras contidas
//na string. Considere que as palavras são separadas por espaços em branco. Caso a string seja vazia, retorne um erro.

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Esta é uma frase de exemplo"
	palavras, err := extrairPalavras(str)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Palavras na string:", palavras)
	}
}

func extrairPalavras(str string) ([]string, error) {
	if str == "" {
		return nil, fmt.Errorf("A string está vazia")
	}

	palavras := strings.Fields(str)
	return palavras, nil
}
