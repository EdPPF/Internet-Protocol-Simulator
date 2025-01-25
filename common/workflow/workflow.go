// Pipeline that sequentially applies each proccessing step on the message modulation.
// Manages types convertions and pass data between methods.
package workflow

import (
	"IP_sim/common/utils"
	"IP_sim/link_layer/error_correction"
	"IP_sim/link_layer/error_detection"
	"IP_sim/link_layer/framing"
	"IP_sim/physical_layer/baseband_modulation"
	"IP_sim/physical_layer/carrier_modulation"
	"fmt"
)

type Workflow struct {
	Message            interface{} // Message muda de tipo entre protocolos e métodos
	Framing            string      // Indica qual protocolo de enquadramento foi selecionado
	ErrorDetection     string      // Indica qual método de detecção de erro foi selecionado
	ErrorCorrection    string      // Indica qual método de correção de erro foi selecionado
	CarrierModulation  string      // Indica qual método de modulação por portadora foi selecionado
	BasebandModulation string      // Indica qual método de modulação em banda base foi selecionado
}

// Inicializa o fluxo de trabalho com a mensagem original
func NewWorkflow(message string) *Workflow {
	return &Workflow{Message: message}
}

func (w *Workflow) Convert(toType string) error {
	switch toType {
	case "string":
		if data, ok := w.Message.([]byte); ok {
			w.Message = string(data)
		} else if data, ok := w.Message.([]int); ok {
			w.Message = utils.FromIntsToString(data)
		} else {
			return fmt.Errorf("conversão inválida: %T para %s", w.Message, toType)
		}

	case "[]byte":
		if data, ok := w.Message.(string); ok {
			w.Message = []byte(data)
		} else if data, ok := w.Message.([]int); ok {
			w.Message = utils.FromIntsToBytes(data)
		} else if data, ok := w.Message.([]uint8); ok {
			w.Message = data
		} else {
			return fmt.Errorf("conversão inválida: %T para %s", w.Message, toType)
		}

	case "[]int":
		if data, ok := w.Message.(string); ok {
			w.Message = utils.ToInt(data)
		} else if data, ok := w.Message.([]byte); ok {
			w.Message = utils.FromBytesToInts(data)
		} else if data, ok := w.Message.([]float64); ok {
			w.Message = utils.FromFloatsToInts(data)
		} else {
			return fmt.Errorf("conversão inválida: %T para %s", w.Message, toType)
		}

	default:
		return fmt.Errorf("conversão inválida: %T para %s", w.Message, toType)
	}
	return nil
}

// ApplyFraming applies the selected framing protocol to the message.
func (w *Workflow) ApplyFraming(method string) error {
	switch method {
	case "Byte Insertion":
		// Converte em []byte se necessário
		if err := w.Convert("[]byte"); err != nil {
			return fmt.Errorf("erro ao converter a mensagem para []byte: %v", err)
		}
		w.Message = framing.EncodeByteInsert(w.Message.([]byte))
		w.Framing = method

	case "Char Count":
		// Converte em []int se necessário
		if err := w.Convert("[]int"); err != nil {
			return fmt.Errorf("erro ao converter a mensagem para []int: %v", err)
		}
		w.Message = framing.CharCountEncode(w.Message.([]int))
		w.Framing = method

	default:
		return fmt.Errorf("protocolo de enquadramento inválido: %s", method)
	}
	return nil
}

// ApplyErrorDetection applies the selected error detection method to the message.
func (w *Workflow) ApplyErrorDetection(method string) error {
	switch method {
	case "Parity Bit":
		// Converte em []int se necessário
		if err := w.Convert("[]int"); err != nil {
			return fmt.Errorf("erro ao converter a mensagem para []int: %v", err)
		}
		w.Message = error_detection.EncodeParity(w.Message.([]int))
		w.ErrorDetection = method

	case "CRC 32":
		// Converte em []byte se necessário
		if err := w.Convert("[]byte"); err != nil {
			return fmt.Errorf("erro ao converter a mensagem para []byte: %v", err)
		}
		w.Message = error_detection.AppendCRC32(w.Message.([]byte))
		w.ErrorDetection = method

	default:
		return fmt.Errorf("método de detecção de erro inválido: %s", method)
	}
	return nil
}

// ApplyErrorCorrection applies the selected error correction method to the message.
func (w *Workflow) ApplyErrorCorrection(method string) error {
	switch method {
	case "Hamming":
		// Converte em []int se necessário
		if err := w.Convert("[]int"); err != nil {
			return fmt.Errorf("erro ao converter a mensagem para []int: %v", err)
		}
		w.Message = error_correction.HammingEncode(w.Message.([]int))
		w.ErrorCorrection = method

	default:
		return fmt.Errorf("método de correção de erro inválido: %s", method)
	}
	return nil
}

// ApplyCarrierModulation applies the selected carrier modulation method to the message.
func (w *Workflow) ApplyCarrierModulation(method string, amplitude float64, f1 float64, f2 float64, frequency float64) error {
	// Converte em []int se necessário
	if err := w.Convert("[]int"); err != nil {
		return fmt.Errorf("erro ao converter a mensagem para []int: %v", err)
	}

	// Extrai os dados da mensagem
	byteStream := w.Message.([]float64)
	// Converte para []int
	bitStream := make([]int, len(byteStream))

	// Aplica os métodos de modulação por portadora
	var modulatedSignal []float64
	switch method {
	case "ASK":
		modulatedSignal = carrier_modulation.AskModulation(amplitude, frequency, bitStream)
	case "FSK":
		modulatedSignal = carrier_modulation.FskModulation(amplitude, f1, f2, bitStream)
	case "8-QAM":
		modulatedSignal = carrier_modulation.Qam8Modulation(amplitude, frequency, bitStream)
	default:
		return fmt.Errorf("unsupported carrier modulation method: %s", method)
	}
	w.Message = modulatedSignal
	w.CarrierModulation = method
	return nil
}

// ApplyBasebandModulation applies the selected baseband modulation method to the message.
func (w *Workflow) ApplyBasebandModulation(method string, V float64) error {
	// Converte em []int se necessário
	/*if err := w.Convert("[]int"); err != nil {
		return fmt.Errorf("erro ao converter a mensagem para []int: %v", err)
	}*/

	// Extrai os dados da mensagem
	bitStream := w.Message.([]int)

	// Aplica os métodos de modulação em banda base
	var modulatedSignal []float64
	switch method {
	case "NRZ Polar":
		modulatedSignal = baseband_modulation.PolarNRZModulation(bitStream, V)
	case "Bipolar":
		modulatedSignal = baseband_modulation.BipolarModulation(bitStream, V)
	case "Manchester":
		modulatedSignal = baseband_modulation.ManchesterModulation(bitStream, V)
	default:
		return fmt.Errorf("método de modulação em banda base inválido: %s", method)
	}
	w.Message = modulatedSignal
	w.BasebandModulation = method
	return nil
}
