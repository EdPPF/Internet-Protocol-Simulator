package server

// Reverses the protocols applied by the client in the correct order.

// Apply the demodulation protocol to the message.
func ApplyDemodulationProtocol(encoded []byte, framing, errorDetection, modulation string) (string, error) {
	demodulated := ReverseModulation(encoded, modulation)
	errorChecked := ReverseErrorDetection(demodulated, errorDetection)
	deframed := ReverseFraming(errorChecked, framing)
	return deframed, nil
}
