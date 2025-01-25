package server

import (
	"IP_sim/common"
	"IP_sim/common/utils"
	"IP_sim/link_layer/error_correction"
	"IP_sim/link_layer/error_detection"
	"IP_sim/link_layer/framing"
	"IP_sim/physical_layer/baseband_modulation"
	"IP_sim/physical_layer/carrier_modulation"
	"fmt"
	"net"
)

func Run() {
	// Inicia o servidor
	listener, err := net.Listen(common.Type, common.Host+":"+common.Port)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor iniciado em " + common.Host + ":" + common.Port + ". Aguardando conexões...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar a conexão:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func ProcessPayload(payload Payload) (string, error) {
	// 1. Undo modulação por Portadora
	floatSignal, err := undoCarrierModulation(payload) // (payload.ModulatedSignal, payload.CarrierFlag)
	if err != nil {
		return "", fmt.Errorf("erro ao desfazer a modulação por portadora: %v", err)
	}

	// 2. Undo modulação por Banda Base
	bitStream, err := undoBasebandModulation(floatSignal, payload.BasebandFlag)
	if err != nil {
		return "", fmt.Errorf("erro ao desfazer a modulação por banda base: %v", err)
	}

	// 3. Undo Hamming
	bitStream, ok := error_correction.HammingDecode(bitStream)
	if ok != 0 {
		return "", fmt.Errorf("erro acusado pelo código de Hamming.\n Bit na posição %v", ok)
	}

	// 4. Verifica com Detecção de erro (CRC ou Parity Bit)
	bitStream, isValid := verifyErrorDetection(bitStream, payload.FramingFlag)
	if !isValid {
		return "", fmt.Errorf("mensagem corrompida")
	}

	// 5. Undo Enquadramento
	bitStream, err = undoFraming(bitStream, payload.FramingFlag)
	if err != nil {
		return "", fmt.Errorf("erro ao desfazer o enquadramento: %v", err)
	}

	// 6. Converte a mensagem para string
	originalMessage := BitsToMessage(bitStream)
	return originalMessage, nil
}

func BitsToMessage(bitStream []int) string {
	// Ensure the bit stream length is a multiple of 8 (1 byte = 8 bits)
	/*if len(bitStream)%8 != 0 {
		panic("Invalid bit stream length: must be a multiple of 8")
	}*/

	var message string
	for i := 0; i < len(bitStream); i += 8 {
		var char byte
		for j := 0; j < 8; j++ {
			char = (char << 1) | byte(bitStream[i+j]) // Shift left and add the current bit
		}
		message += string(char) // Convert the byte to its ASCII character and append to the message
	}

	return message
}

func undoFraming(bitStream []int, framingFlag string) ([]int, error) {
	switch framingFlag {
	case "Char Count":
		return framing.CharCountDecode(bitStream)
	case "Byte Insertion":
		// Converte bitStream para bytes
		byteStream := utils.FromIntsToBytes(bitStream)
		return framing.DecodeByteInsert(byteStream)
	default:
		return nil, fmt.Errorf("método de enquadramento inválido: %s", framingFlag)
	}
}

func verifyErrorDetection(bitStream []int, framingFlag string) ([]int, bool) {
	switch framingFlag {
	case "Parity Bit":
		return error_detection.DecodeParity(bitStream)
	case "CRC 32":
		// Converte bitStream para bytes
		byteStream := utils.FromIntsToBytes(bitStream)
		return error_detection.VerifyCRC32(byteStream)
	default:
		return nil, false
	}
}

func undoBasebandModulation(floatSignal []float64, method string) ([]int, error) {
	switch method {
	case "NRZ Polar":
		return baseband_modulation.PolarNRZDemodulation(floatSignal), nil
	case "Bipolar":
		return baseband_modulation.BipolarDemodulation(floatSignal), nil
	case "Manchester":
		return baseband_modulation.ManchesterDemodulation(floatSignal), nil
	default:
		return nil, fmt.Errorf("método de modulação em banda base inválido: %s", method)
	}
}

func undoCarrierModulation(payload Payload) ([]float64, error) {
	switch payload.CarrierFlag {
	case "ASK":
		return carrier_modulation.AskDemodulation(payload.ModulatedSignal), nil
	case "FSK":
		return carrier_modulation.FskDemodulation(payload.ModulatedSignal, payload.Amplitude, payload.F1, payload.F2), nil
	case "QAM 8":
		return carrier_modulation.Qam8Demodulation(payload.ModulatedSignal, payload.Amplitude, payload.Frequency), nil
	default:
		return nil, fmt.Errorf("método de modulação por portadora inválido: %s", payload.CarrierFlag)
	}
}
