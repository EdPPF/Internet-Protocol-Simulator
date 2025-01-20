package framing

import (
	"errors"
	"fmt"
)

// Escreve um cabeçalho de tamanho no início do array de dados.
func EncodeCharCount(data []int) []int {
	x := len(data)
	data = append([]int{x}, data...)
	return data
}

func EncodeCharCountWrapper(data interface{}) (interface{}, error) {
	d, ok := data.([]int)
	if !ok {
		return nil, fmt.Errorf("invalid input type for EncodeCharCount")
	}
	return EncodeCharCount(d), nil
}

// Remove o cabeçalho de tamanho do array de dados.
func DecodeCharCount(data []int) ([]int, error) {
	if data[0] == len(data)-1 {
		return append(data[:0], data[0+1:]...), nil
	} else {
		return data, errors.New("transmission failed")
	}
}

func DecodeCharCountWrapper(data interface{}) (interface{}, error) {
	d, ok := data.([]int)
	if !ok {
		return nil, fmt.Errorf("invalid input type for DecodeCharCount")
	}
	return DecodeCharCount(d)
}
