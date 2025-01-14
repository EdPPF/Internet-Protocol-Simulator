package main

import (
	"fmt"
)

// Codifica a mensagem recebida utilizando paridade par.
func encodeParity(data []int) []int {
	parity := 0
	for _, bit := range data { // range pega o índice e o valor, por isso _
		parity ^= bit
	}
	return append(data, parity)
}

// Decodifica a mensagem recebida utilizando paridade par.
func decodeParity(dataWithParity []int) ([]int, bool) {
	// Mensagem vazia
	if len(dataWithParity) == 0 {
		return nil, false
	}

	// Calcula a paridade
	parity := 0
	for _, bit := range dataWithParity {
		parity ^= bit
	}

	// parity==0 indica que a mensagem foi recebida corretamente
	return dataWithParity[:len(dataWithParity)-1], parity == 0
}

func main() {
	// Data original
	data := []int{1, 0, 1, 1, 0, 0, 1}
	fmt.Println("Data original:", data)

	// Encode
	dataWithParity := encodeParity(data)
	fmt.Println("Data with parity:", dataWithParity)

	// Decode
	decoded, ok := decodeParity(dataWithParity)
	fmt.Println("Data:", decoded, "Parity check:", ok, " -> OK")

	// Com erro
	dataWithParity[2] = 0
	decoded, ok = decodeParity(dataWithParity)
	fmt.Println("Data:", decoded, "Parity check:", ok, "-> ERRO")
}
