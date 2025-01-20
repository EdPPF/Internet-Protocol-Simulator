// Applies the selected framing, error detection, and modulation protocols to the message.
// Combines protocol functions from `link_layer/` and `physical_layer/`.
package client

func ApplyProtocols(message string, framing, errorDetection, modulation string) ([]byte, error) {
	// Apply framing protocol
	framedMessage, err := ApplyFramingProtocol(message, framing)
	if err != nil {
		return nil, err
	}

	// Apply error detection protocol
	messageWithChecksum, err := ApplyErrorDetectionProtocol(framedMessage, errorDetection)
	if err != nil {
		return nil, err
	}

	// Apply modulation protocol
	modulatedMessage, err := ApplyModulationProtocol(messageWithChecksum, modulation)
	if err != nil {
		return nil, err
	}

	return modulatedMessage, nil
}
