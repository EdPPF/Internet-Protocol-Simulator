// Utilitários para codificação e decodificação vários tipos de dado.
package utils

import (
	"fmt"
)

// ToBytes converte uma string em um slice de bytes.
func ToBytes(s string) []byte {
	return []byte(s)
}

// ToString converte um slice de bytes em uma string.
func ToString(b []byte) string {
	return string(b)
}

func StringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%.8b", binString, c)
	}
	return
}

// ToInt converte uma string de bits em uma slice de inteiros.
func ToInt(s string) []int {
	var ints []int
	for _, c := range s {
		if c == '0' {
			ints = append(ints, 0)
		} else {
			ints = append(ints, 1)
		}
	}
	return ints
}

// FromIntsToString converts a slice of ints to a binary string.
func FromIntsToString(data []int) string {
	result := ""
	for _, bit := range data {
		if bit == 1 {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}

// FromIntsToBytes converts a slice of ints to a slice of bytes.
func FromIntsToBytes(data []int) []byte {
	var result []byte
	var currentByte byte
	bitCount := 0

	for _, bit := range data {
		currentByte = (currentByte << 1) | byte(bit) // Shift left and add the bit
		bitCount++

		if bitCount == 8 { // When 8 bits are accumulated, add the byte to the result
			result = append(result, currentByte)
			currentByte = 0
			bitCount = 0
		}
	}

	// Handle any remaining bits (pad with 0s to complete the last byte)
	if bitCount > 0 {
		currentByte = currentByte << (8 - bitCount) // Pad with 0s
		result = append(result, currentByte)
	}

	return result
}

/* func FromIntsToBytes(data []int) []byte {
	result := make([]byte, len(data))
	for i, bit := range data {
		result[i] = byte(bit)
	}
	return result
}
*/

// FromBytesToInts converts a slice of bytes to a slice of ints.
func FromBytesToInts(data []byte) []int {
	var bits []int
	for _, b := range data {
		for i := 7; i >= 0; i-- { // Extract each bit (MSB to LSB)
			bits = append(bits, int((b>>i)&1))
		}
	}
	return bits
}

/*func FromBytesToInts(data []byte) []int {
	result := make([]int, len(data))
	for i, b := range data {
		result[i] = int(b)
	}
	return result
}*/

// FromFloatsToInts converts a slice of float64 values to a slice of ints.
func FromFloatsToInts(data []float64) []int {
	var result []int
	for _, value := range data {
		result = append(result, int(value))
	}
	return result
}

// Float64ToInt converts a slice of float64 values to a slice of ints based on a threshold.
// Utilizada para passar o sinal demodulado da banda base para demodulação do código de Hmmaing.
// threshold é escolhido baseado no método utilizado na modulação banda base.
func Float64ToInt(signal []float64, threshold float64) []int {
	var bitStream []int
	for _, value := range signal {
		if value > threshold {
			bitStream = append(bitStream, 1)
		} else {
			bitStream = append(bitStream, 0)
		}
	}
	return bitStream
}
