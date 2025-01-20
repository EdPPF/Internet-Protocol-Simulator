package main

import (
	"errors"
	"fmt"
)

// Escreve um cabeçalho de tamanho no início do array de dados.
func CharCountEncode(data []int) []int {
	x := len(data)
	data = append([]int{x}, data...)
	return data
}

// Remove o cabeçalho de tamanho do array de dados.
func CharCountDecode(data []int) ([]int, error) {
	if data[0] == len(data)-1 {
		return append(data[:0], data[0+1:]...), nil
	} else {
		return data, errors.New("transmission failed")
	}
}

func main() {

	// Example input data
	data := []int{1, 0, 1, 1, 0}
	fmt.Println("Original data:", data)

	encoded := CharCountEncode(data)
	fmt.Println("Encoded with char. count:", encoded)
	fmt.Println("Data header (size):", encoded[0])

	decoded, err := CharCountDecode(encoded)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Decoded data:", decoded)
	}
}
