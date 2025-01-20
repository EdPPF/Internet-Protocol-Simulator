package error_detection

// package error_detection

import (
	"IP_sim/common"
	"fmt"
)

// CRC32Table stores the precomputed CRC32 lookup table
var CRC32Table [256]uint32

// Essa função é usada para inverter os bytes de um valor de 32 bits?
func reverseBytes32(value uint32) []int {
	result := []int{
		int((value >> 24) & 0xFF),
		int((value >> 16) & 0xFF),
		int((value >> 8) & 0xFF),
		int(value & 0xFF),
	}
	return result
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

// ComputeCRC32 calculates the CRC32 checksum for the given data slice
func ComputeCRC32(data []int) []int {
	crc := uint32(0xFFFFFFFF) // Initialize CRC with all bits set according to IEEE 802.1
	for _, b := range data {
		tableIndex := byte(crc>>24) ^ byte(b)
		crc = (crc << 8) ^ CRC32Table[tableIndex]
	}
	return reverseBytes32(^crc) // Final XOR to invert the bits
}

func ComputeCRC32Wrapper(input interface{}) (interface{}, error) {
	data, ok := input.([]int)
	if !ok {
		return nil, fmt.Errorf("invalid input type for ComputeCRC32")
	}

	result := ComputeCRC32(data)
	return result, nil
}

/*
func main() {
	data := []byte("hello")
	// data := []byte("123456789")
	crc := ComputeCRC32(data) // 1931653D | 181989fc
	fmt.Printf("CRC32 Checksum (Big-Endian): %08X\n", crc)
}
*/
