package error_detection

// package error_detection

import (
	"IP_sim/common"
	"IP_sim/common/utils"
	"encoding/binary"
)

// CRC32Table stores the precomputed CRC32 lookup table
var CRC32Table [256]uint32

// Inverte a ordem dos bytes de um uint32 para refletir a arquitetura little-endian.
// ! Não está funcionando!
func reverseBytes32(value uint32) uint32 {
	return (value>>24)&0xFF | // Move byte 0 to byte 3
		((value>>8)&0xFF)<<8 | // Move byte 1 to byte 2
		((value<<8)&0xFF)<<16 | // Move byte 2 to byte 1
		(value<<24)&0xFF000000 // Move byte 3 to byte 0
}

// init precomputes the CRC32 table using the polynomial from common/constants.go
func init() {
	polynomial := common.CRC_polynomial
	for i := 0; i < 256; i++ {
		crc := uint32(i) << 24
		for j := 0; j < 8; j++ {
			if crc&0x80000000 != 0 {
				crc = (crc << 1) ^ polynomial
			} else {
				crc <<= 1
			}
		}
		CRC32Table[i] = crc
	}
}

// computeCRC32 calculates the CRC32 checksum for the given data slice
func computeCRC32(data []byte) uint32 {
	crc := uint32(0xFFFFFFFF) // Initialize CRC with all bits set according to IEEE 802.1
	for _, b := range data {
		tableIndex := byte(crc>>24) ^ b
		crc = (crc << 8) ^ CRC32Table[tableIndex]
	}
	return ^crc // Final XOR to invert the bits
}

// VerifyCRC32 validates the integrity of the received data using the CRC32 checksum.
// Returns true if the data is valid, false otherwise.
func VerifyCRC32(dataWithCRC []byte) ([]int, bool) {
	// Split the received data into the data portion and the CRC portion
	if len(dataWithCRC) < 4 {
		// Message is too short to contain both data and CRC
		return nil, false
	}

	data := dataWithCRC[:len(dataWithCRC)-4]                                 // Extract the data portion
	receivedCRC := binary.BigEndian.Uint32(dataWithCRC[len(dataWithCRC)-4:]) // Extract the CRC

	// Recompute the CRC from the data
	computedCRC := computeCRC32(data)

	// Compare the received CRC with the recomputed CRC
	isValid := receivedCRC == computedCRC

	// Convert the data to a slice of bits
	bits := utils.FromBytesToInts(data)
	return bits, isValid
}

// AppendCRC32 calculates the CRC32 checksum and append it to the given data slice.
func AppendCRC32(data []byte) []byte {
	crc := computeCRC32(data)
	crcBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(crcBytes, crc)
	return append(data, crcBytes...)
}
