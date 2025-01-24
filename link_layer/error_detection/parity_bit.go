package error_detection

// Codifica a mensagem recebida utilizando paridade par.
func EncodeParity(data []int) []int {
	parity := 0
	for _, bit := range data { // range pega o índice e o valor, por isso _
		parity ^= bit
	}
	return append(data, parity)
}

// Decodifica a mensagem recebida utilizando paridade par.
func DecodeParity(dataWithParity []int) ([]int, bool) {
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
