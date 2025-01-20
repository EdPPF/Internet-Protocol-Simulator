package communication

import (
	"IP_sim/link_layer/error_correction"
	"IP_sim/link_layer/error_detection"
	"IP_sim/link_layer/framing"
	"IP_sim/physical_layer/baseband_modulation"
	"IP_sim/physical_layer/carrier_modulation"
)

// Estrutura de um protocolo de comunicação.
type Protocol struct {
	Name    string                                 // Identifica o protocolo.
	Apply   func(interface{}) (interface{}, error) // Aplica o protocolo.
	Reverse func(interface{}) (interface{}, error) // Reverte o protocolo.
}

var SupportedProtocols = map[string]Protocol{
	// Error Correction
	"Hamming": {Name: "Hamming", Apply: error_correction.HammingEncodeWrapper, Reverse: error_correction.HammingDecodeWrapper},
	// Error Detection
	"CRC":    {Name: "CRC", Apply: error_detection.ComputeCRC32Wrapper, Reverse: nil},
	"Parity": {Name: "Parity", Apply: error_detection.EncodeParityWrapper, Reverse: error_detection.DecodeParityWrapper},
	// Modulation
	// Baseband
	"PolarNRZ":   {Name: "Polar NRZ", Apply: baseband_modulation.PolarNRZModulationWrapper, Reverse: baseband_modulation.PolarNRZDemodulationWrapper},
	"Manchester": {Name: "Manchester", Apply: baseband_modulation.ManchesterModulationWrapper, Reverse: baseband_modulation.ManchesterDemodulationWrapper},
	"Bipolar":    {Name: "Bipolar", Apply: baseband_modulation.BipolarModulationWrapper, Reverse: baseband_modulation.BipolarDemodulationWrapper},
	// Carrier
	"ASK":   {Name: "ASK", Apply: carrier_modulation.AskModulationWrapper, Reverse: carrier_modulation.AskDemodulationWrapper},
	"FSK":   {Name: "FSK", Apply: carrier_modulation.FskModulationWrapper, Reverse: carrier_modulation.FskDemodulationWrapper},
	"QAM-8": {Name: "QAM-8", Apply: carrier_modulation.Qam8ModulationWrapper, Reverse: carrier_modulation.Qam8DemodulationWrapper},
	// Framing
	"CharCount":  {Name: "CharCount", Apply: framing.EncodeCharCountWrapper, Reverse: framing.DecodeCharCountWrapper},
	"ByteInsert": {Name: "ByteInsert", Apply: framing.EncodeByteInsertWrapper, Reverse: framing.DecodeByteInsertWrapper},
}
