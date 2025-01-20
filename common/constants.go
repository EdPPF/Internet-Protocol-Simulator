// Constantes compartilhadas
package common

const (
	// 0.0.0.0 listens on all network interfaces
	Host = "0.0.0.0" // "localhost"
	Port = "8080"
	Type = "tcp"
	// IEEE 802 CRC-32 polynomial
	CRC_polynomial         uint32 = 0x04C11DB7
	reverse_CRC_polynomial uint32 = 0xEDB88320
	// ASCII Constants for framing
	STX = byte(0x02) // Start of Text
	ETX = byte(0x03) // End of Text
	ESC = byte(0x1B) // Escape
)
