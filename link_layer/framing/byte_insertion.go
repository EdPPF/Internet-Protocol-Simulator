package framing

import (
	"IP_sim/common"
	"bytes"
	"errors"
)

// Insere bytes ou chars no início e fim do array de dados.
// Cada quadro começa e termina com um byte especial, chamado de byte de flag.
// O enquadramento é feitoc por adição de caracteres ASCII.
func EncodeByteInsert(data []byte) []byte {
	var framed bytes.Buffer

	// Adiciona o início do quadro
	framed.WriteByte(common.STX)

	for _, b := range data {
		// Se o byte é um byte de flag ou um byte de escape, adicione um byte de escape
		if b == common.STX || b == common.ETX || b == common.ESC {
			framed.WriteByte(common.ESC) // Add escape character
		}
		framed.WriteByte(b) // Escreve o byte
	}

	// Adiciona o fim do quadro
	framed.WriteByte(common.ETX)

	return framed.Bytes()
}

func EncodeByteInsertWrapper(data interface{}) (interface{}, error) {
	d, ok := data.([]byte)
	if !ok {
		return nil, errors.New("invalid input type for EncodeByteInsert")
	}
	return EncodeByteInsert(d), nil
}

// Remove bytes do início e fim do array de dados.
func DencodeByteInsert(framedData []byte) ([]byte, error) {
	// Verifica se o quadro é válido
	if len(framedData) < 2 || framedData[0] != common.STX || framedData[len(framedData)-1] != common.ETX {
		return nil, errors.New("invalid frame")
	}

	var unframed bytes.Buffer
	escaping := false

	// Pula STX e ETX e desfaz o enquadramento
	for _, b := range framedData[1 : len(framedData)-1] {
		if escaping {
			// Se o byte anterior é um byte de escape, interpreta o byte atual literalmente
			unframed.WriteByte(b)
			escaping = false
			continue
		}

		if b == common.ESC {
			// Se o byte atual é um byte de escape, pule-o e marque o próximo byte para ser interpretado literalmente
			escaping = true
			continue
		}

		// Byte normal
		unframed.WriteByte(b)
	}

	if escaping {
		return nil, errors.New("malformed frame (dangling ESC)")
	}

	return unframed.Bytes(), nil
}

func DecodeByteInsertWrapper(data interface{}) (interface{}, error) {
	d, ok := data.([]byte)
	if !ok {
		return nil, errors.New("invalid input type for DecodeByteInsert")
	}
	return DencodeByteInsert(d)
}
