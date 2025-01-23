// Utilitários para codificação e decodificação de strings
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
