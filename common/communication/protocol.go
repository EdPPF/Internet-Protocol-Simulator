package communication

import (
	"IP_sim/link_layer/error_correction"
	"IP_sim/link_layer/error_detection"
)

// Estrutura de um protocolo de comunicação.
type Protocol struct {
	Name    string                   // Identifica o protocolo.
	Apply   func([]int) []int        // Aplica o protocol.
	Reverse func([]int) ([]int, int) // Reverte o protocolo.
}

var SupportedProtocols = map[string]Protocol{
	// Error Correction
	"Hamming": Protocol{Name: "Hamming", Apply: error_correction.HammingEncode, Reverse: error_correction.HammingDecode},
	// Error Detection
	"CRC":    Protocol{Name: "CRC", Apply: error_detection.ComputeCRC32, Reverse: nil},
	"Parity": Protocol{Name: "Parity", Apply: error_detection.EncodeParity, Reverse: error_detection.DecodeParity},
	// Modulation
	// Baseband
	"NRZ":        Protocol{Name: "NRZ", Apply: nil, Reverse: nil},
	"Manchester": Protocol{Name: "Manchester", Apply: nil, Reverse: nil},
	"Bipolar":    Protocol{Name: "Bipolar", Apply: nil, Reverse: nil},
	// Carrier
	"ASK":   Protocol{Name: "ASK", Apply: nil, Reverse: nil},
	"FSK":   Protocol{Name: "FSK", Apply: nil, Reverse: nil},
	"QAM-8": Protocol{Name: "QAM-8", Apply: nil, Reverse: nil},
}
