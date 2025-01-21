package framing

import (
	"IP_sim/common/utils"
	"errors"
	"fmt"
	"log"
)

// Escreve um cabeçalho de tamanho no início do array de dados.
func EncodeCharCount(data string) []int {
	ints, err := utils.ParseBinaryString(data)
	if err != nil {
		log.Fatalf("Error parsing binary string: %v", err)
	}
	x := len(ints)
	ints = append([]int{x}, ints...)
	return ints
}

func EncodeCharCountWrapper(data interface{}) (interface{}, error) {
	d, ok := data.(string)
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
