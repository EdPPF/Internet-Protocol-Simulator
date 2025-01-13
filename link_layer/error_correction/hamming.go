package main

import (
	"fmt"
	"math"
)

// HammingEncode encodes the given data using Hamming(7,4) code
func HammingEncode(data []int) []int {
	m := len(data) // Número de bits na mensagem original
	r := calculateRedundantBits(m)

	// Mensagem codificada com espaço para os bits de paridade
	encoded := make([]int, m+r+1) // +1 for 1-based indexing

	// Insere os bits de dados nas posições corretas
	j := 0
	for i := 1; i < len(encoded); i++ {
		if isPowerOfTwo(i) {
			continue // Pula as posições de paridade, que são as potências de 2
		}
		encoded[i] = data[j]
		j++
	}

	// Calcula os bits de paridade
	for i := 0; i < r; i++ {
		parityPos := int(math.Pow(2, float64(i)))
		encoded[parityPos] = calculateParity(encoded, parityPos)
	}

	return encoded[1:] // Remove the 0th index for 1-based indexing
}

// Decodifica a mensagem recebida e corrige um (1, uno, one, 一) erro, se houver
func HammingDecode(received []int) ([]int, int) {
	n := len(received)
	r := int(math.Log2(float64(n + 1)))

	// Calcula a posição do bit de paridade e o valor do bit de paridade
	syndrome := 0
	for i := 0; i < r; i++ {
		parityPos := int(math.Pow(2, float64(i)))
		parityValue := calculateParity(received, parityPos)
		if parityValue != 0 {
			syndrome += parityPos
		}
	}

	// Corrige o erro, se houver
	if syndrome != 0 && syndrome <= n {
		received[syndrome] ^= 1 // Flip the erroneous bit
	}

	// Extrai os bits de dados da mensagem recebida
	data := []int{}
	for i := 1; i <= n; i++ {
		if !isPowerOfTwo(i) {
			data = append(data, received[i-1])
		}
	}

	return data, syndrome
}

// Determina o número de bits de paridade de acordo com a regra: 2^r >= m+r+1
func calculateRedundantBits(dataBits int) int {
	for r := 1; ; r++ {
		if int(math.Pow(2, float64(r))) >= dataBits+r+1 {
			return r
		}
	}
}

// Calcula o valor do bit de paridade para uma posição específica
func calculateParity(bits []int, parityPos int) int {
	parity := 0
	for i := 1; i < len(bits); i++ {
		if i&parityPos != 0 {
			parity ^= bits[i] // parity = parity XOR bit
		}
	}
	return parity
}

// Nome autoexplicativo...
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func main() {
	// Example usage
	data := []int{1, 0, 1, 1} // 4 data bits
	fmt.Println("Original Data:", data)

	// Encode data
	encoded := HammingEncode(data)
	fmt.Println("Encoded Data:", encoded)

	// Introduce an error for testing
	encoded[3] ^= 1 // Flip a bit
	fmt.Println("Received Data with Error:", encoded)

	// Decode data
	decoded, syndrome := HammingDecode(encoded)
	fmt.Println("Decoded Data:", decoded)
	if syndrome != 0 {
		fmt.Printf("Error detected and corrected at position: %d\n", syndrome)
	} else {
		fmt.Println("No error detected.")
	}
}
