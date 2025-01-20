package error_correction

import (
	"fmt"
	"math"
)

// HammingEncode encodes the given data using Hamming code
func HammingEncode(data []int) []int {
	m := len(data) // Número de bits na mensagem original
	r := calculateRedundantBits(m)

	// Mensagem codificada com espaço para os bits de paridade
	encoded := make([]int, m+r)

	// Insere os bits de dados nas posições corretas
	j := 0
	for i := 0; i < len(encoded); i++ {
		if isPowerOfTwo(i + 1) { // if (1 << i) & 1 != 0 ???
			continue // Pula as posições de paridade, que são as potências de 2
		}
		encoded[i] = data[j]
		j++
	}

	// Calcula os bits de paridade
	for i := 0; i < r; i++ {
		parityPos := (1 << i) - 1 // 0-based parity positions: 0, 1, 3, 7...
		encoded[parityPos] = calculateParity(encoded, parityPos)
	}

	return encoded
}

func HammingEncodeWrapper(input interface{}) (interface{}, error) {
	data, ok := input.([]int)
	if !ok {
		return nil, fmt.Errorf("invalid input type for HammingEncode")
	}

	result := HammingEncode(data)
	return result, nil
}

// Decodifica a mensagem recebida e corrige um (1, uno, one, 一) erro, se houver
func HammingDecode(received []int) ([]int, int) {
	n := len(received)
	r := int(math.Ceil(math.Log2(float64(n + 1))))

	// Calcula a posição do bit de paridade e o valor do bit de paridade
	syndrome := 0
	for i := 0; i < r; i++ {
		parityPos := (1 << i) - 1 // 0-based parity positions: 0, 1, 3, 7...
		parityValue := calculateParity(received, parityPos)
		if parityValue != 0 {
			syndrome += (parityPos + 1) // 1-based position for syndrome
		}
	}

	// Corrige o erro, se houver
	if syndrome != 0 && syndrome <= n {
		received[syndrome-1] ^= 1 // Flip the erroneous bit
	}

	// Extrai os bits de dados da mensagem recebida
	data := []int{}
	for i := 0; i < n; i++ {
		if !isPowerOfTwo(i + 1) {
			data = append(data, received[i])
		}
	}

	return data, syndrome
}

func HammingDecodeWrapper(input interface{}) (interface{}, error) {
	received, ok := input.([]int)
	if !ok {
		return nil, fmt.Errorf("invalid input type for HammingDecode")
	}

	data, syndrome := HammingDecode(received)
	return struct {
		Data     []int
		Syndrome int
	}{Data: data, Syndrome: syndrome}, nil
}

// Determina o número de bits de paridade de acordo com a regra: 2^r >= m+r+1
func calculateRedundantBits(dataBits int) int {
	for r := 0; ; r++ {
		if (1 << r) >= dataBits+r+1 {
			return r
		}
	}
}

// Calcula o valor do bit de paridade para uma posição específica
func calculateParity(bits []int, parityPos int) int {
	parity := 0
	for i := 0; i < len(bits); i++ {
		// Check if the bit at position `i` contributes to the parity bit at `parityPos`
		if (i+1)&(parityPos+1) != 0 { // Use (i+1) for 1-based parity logic
			parity ^= bits[i]
		}
	}
	return parity
}

// Verifica se um número é potência de 2...
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

/*
func main() {
	// data := []int{1, 1, 0, 1, 0, 0, 1} // Data original
	data := []int{1, 1, 1, 0, 1, 1}
	fmt.Println("Original Data:", data)

	// Encode
	// encoded := HammingEncode(data) // 01101011001
	encoded := HammingEncode(data)
	fmt.Println("Encoded Data (1 1 1 0 1 1 0 0 1 1):", encoded)

	// Introduz erro
	encoded[4] ^= 1 // Flip a bit
	fmt.Println("Received Data with Error (1 1 1 0 0 1 0 0 1 1):", encoded)

	// Decode
	decoded, syndrome := HammingDecode(encoded)
	fmt.Println("Syndrome:", syndrome)

	fmt.Println("Decoded Data:", decoded)
	if syndrome != 0 {
		fmt.Printf("Error detected and corrected at position: %d\n", syndrome)
	} else {
		fmt.Println("No error detected.")
	}
}
*/
