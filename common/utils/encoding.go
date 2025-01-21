// Utilitários para codificação e decodificação de strings
package utils

import (
	"errors"
	"fmt"
)

func ToBytes(s string) []byte {
	return []byte(s)
}

func ToString(b []byte) string {
	return string(b)
}

func BytesToInt(b []byte) []int {
	var binaryRepresentation []int

	for _, byteVal := range b {
		// Process each byte and extract its 8 bits
		for i := 7; i >= 0; i-- {
			bit := (byteVal >> i) & 1
			binaryRepresentation = append(binaryRepresentation, int(bit))
		}
	}

	return binaryRepresentation
}

func StringToBinary(s string) []int {
	var binaryRepresentation []int

	for _, char := range s {
		// Convert character to its ASCII value (int32) and process 8 bits
		for i := 7; i >= 0; i-- {
			bit := (char >> i) & 1
			binaryRepresentation = append(binaryRepresentation, int(bit))
		}
	}

	return binaryRepresentation
}
func ToStringBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}

func ParseBinaryString(data string) ([]int, error) {
	var result []int
	for _, char := range data {
		// Convert each character to its ASCII value
		binary := fmt.Sprintf("%08b", char) // Convert to an 8-bit binary string
		for _, bit := range binary {
			if bit != '0' && bit != '1' {
				return nil, errors.New("unexpected error converting character to binary")
			}
			result = append(result, int(bit-'0')) // Convert '0'/'1' to 0/1
		}
	}
	return result, nil
}
