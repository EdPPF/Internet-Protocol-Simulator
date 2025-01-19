package communication

import (
	"IP_sim/link_layer/error_correction"
	"IP_sim/link_layer/error_detection"
	"IP_sim/physical_layer/baseband_modulation"
	"IP_sim/physical_layer/carrier_modulation"
)

// Estrutura de um protocolo de comunicação.
type Protocol struct {
	Name string // Identifica o protocolo.
	// Apply   func([]int) []int        // Aplica o protocolo.
	// Reverse func([]int) ([]int, int) // Reverte o protocolo.
	Apply   func(interface{}) (interface{}, error) // Aplica o protocolo.
	Reverse func(interface{}) (interface{}, error) // Reverte o protocolo.
}

var SupportedProtocols = map[string]Protocol{
	// Error Correction
	"Hamming": Protocol{Name: "Hamming", Apply: error_correction.HammingEncodeWrapper, Reverse: error_correction.HammingDecodeWrapper},
	// Error Detection
	"CRC":    Protocol{Name: "CRC", Apply: error_detection.ComputeCRC32Wrapper, Reverse: nil},
	"Parity": Protocol{Name: "Parity", Apply: error_detection.EncodeParityWrapper, Reverse: error_detection.DecodeParityWrapper},
	// Modulation
	// Baseband
	"PolarNRZ":   Protocol{Name: "Polar NRZ", Apply: baseband_modulation.PolarNRZModulationWrapper, Reverse: baseband_modulation.PolarNRZDemodulationWrapper},
	"Manchester": Protocol{Name: "Manchester", Apply: baseband_modulation.ManchesterModulationWrapper, Reverse: baseband_modulation.ManchesterDemodulationWrapper},
	"Bipolar":    Protocol{Name: "Bipolar", Apply: baseband_modulation.BipolarModulationWrapper, Reverse: baseband_modulation.BipolarDemodulationWrapper},
	// Carrier
	"ASK":   Protocol{Name: "ASK", Apply: carrier_modulation.AskModulationWrapper, Reverse: carrier_modulation.AskDemodulationWrapper},
	"FSK":   Protocol{Name: "FSK", Apply: carrier_modulation.FskModulationWrapper, Reverse: carrier_modulation.FskDemodulationWrapper},
	"QAM-8": Protocol{Name: "QAM-8", Apply: carrier_modulation.Qam8ModulationWrapper, Reverse: carrier_modulation.Qam8DemodulationWrapper},
}
