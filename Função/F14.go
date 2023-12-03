//Escreva uma função que receba dois slices de inteiros como parâmetros e retorne um novo
//slice contendo apenas os números presentes nos dois slices. Caso um dos slices esteja vazio, retorne um erro.

package main

import (
	"errors"
	"fmt"
)

func intersection(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("Um dos slices está vazio")
	}

	numbersMap := make(map[int]bool)
	for _, num := range slice1 {
		numbersMap[num] = true
	}

	var result []int
	for _, num := range slice2 {
		if numbersMap[num] {
			result = append(result, num)
		}
	}

	return result, nil
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{3, 4, 5, 6, 7}
	intersectedSlice, err := intersection(slice1, slice2)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Interseção dos slices:", intersectedSlice)
	}
}
