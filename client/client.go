package client

import (
	"IP_sim/common/utils"
	"IP_sim/link_layer/framing"
	"IP_sim/physical_layer/baseband_modulation"
	"IP_sim/physical_layer/carrier_modulation"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var FramingProtocols = []string{"Byte Insertion", "Char Count"}
var CarrierModulation = []string{"ASK", "FSK", "QAM 8"}
var BasebandModulation = []string{"NRZ Polar", "Bipolar", "Manchester"}

func Run() {
	fmt.Println("Cliente iniciado.")

	// User input = mensagem a ser enviada ao server
	fmt.Println("Digite a mensagem a ser enviada:")
	reader := bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')
	message = strings.TrimSuffix(message, "\n")

	// Transforma a mensagem em uma string de bits
	bits := utils.StringToBin(message)
	fmt.Println("- Mensagem original:", message)
	fmt.Println("+ Mensagem em bits:", bits)
	// Transforma a string de bits em um array de inteiros
	data := utils.ToInt(bits)

	// User input: Seleciona protocolo de enquadramento (byte_insertion ou char_count)
	var frameMethod int
	var framingFlag string // Flag para indicar qual protocolo de enquadramento foi selecionado
	fmt.Println("Selecione um protocolo de enquadramento:")
	for {
		for i, protocol := range FramingProtocols {
			fmt.Printf("%d. %s\n", i+1, protocol)
		}
		fmt.Scan(&frameMethod)
		if frameMethod == 1 || frameMethod == 2 {
			break
		}
		fmt.Println("Protocolo inválido. Por favor, selecione um dos protocolos abaixo (1 | 2).")
	}
	// Chama as funções de enquadramento
	switch frameMethod {
	case 1: // Byte Insertion
		framingFlag = "Byte Insertion"
		fmt.Println("Aplicando Byte Insertion...")
		byts := utils.ToBytes(message)
		fmt.Println("Original Data (string):", message)
		fmt.Println("Original Data (bytes ):", byts)
		framed := framing.EncodeByteInsert(byts)
		fmt.Println("Framed Data (bytes )  :", framed)
		fmt.Printf("Framed Data (string)  : %q\n", utils.ToString(framed)) // %q permite que os bytes de escape sejam exibidos no print
	case 2: // Char Count
		framingFlag = "Char Count"
		fmt.Println("Aplicando Char Count...")
		fmt.Println("Original Data:", data)
		framed := framing.CharCountEncode(data)
		fmt.Println("Framed Data:", framed)
		fmt.Println("Data header (size):", framed[0])
	default:
		fmt.Println("Protocolo de enquadramento inválido!")
	}

	// User input: Seleciona modulação digital (carrier modulation)
	var carrier int
	var carrierFlag string // Flag para indicar qual método de modulação por portadora foi selecionado
	fmt.Println("Selecione um método de modulação por portadora:")
	for {
		for i, mod := range CarrierModulation {
			fmt.Printf("%d. %s\n", i+1, mod)
		}
		fmt.Scan(&carrier)
		if carrier == 1 || carrier == 2 || carrier == 3 {
			break
		}
		fmt.Println("Método de modulação inválido. Por favor, selecione um dos métodos abaixo (1 | 2 | 3).")
	}
	// Chama as funções de modulação por portadora
	switch carrier {
	case 1: // ASK
		carrierFlag = "ASK"
		var A, F float64
		fmt.Println("Digite o valor de Amplitude:")
		fmt.Scan(&A)
		fmt.Println("Digite o valor de Frequência:")
		fmt.Scan(&F)
		fmt.Println("Aplicando ASK...")
		modulated := carrier_modulation.AskModulation(A, F, data)
		fmt.Println("Modulated Data:", modulated)

	case 2: // FSK
		carrierFlag = "FSK"
		var A, F1, F2 float64
		fmt.Println("Digite o valor de Amplitude:")
		fmt.Scan(&A)
		fmt.Println("Digite o valor de Frequência 1:")
		fmt.Scan(&F1)
		fmt.Println("Digite o valor de Frequência 2:")
		fmt.Scan(&F2)
		fmt.Println("Aplicando FSK...")
		modulated := carrier_modulation.FskModulation(A, F1, F2, data)
		fmt.Println("Modulated Data:", modulated)

	case 3: // QAM 8
		carrierFlag = "QAM 8"
		var A, F float64
		fmt.Println("Digite o valor de Amplitude:")
		fmt.Scan(&A)
		fmt.Println("Digite o valor de Frequência:")
		fmt.Scan(&F)
		fmt.Println("Aplicando QAM 8...")
		modulated := carrier_modulation.Qam8Modulation(A, F, data)
		fmt.Println("Modulated Data:", modulated)

	default:
		fmt.Println("Método de modulação inválido!")
	}

	// User input: Seleciona modulação analógica (baseband modulation)
	var baseband int
	var basebandFlag string // Flag para indicar qual método de modulação em banda base foi selecionado
	fmt.Println("Selecione um método de modulação em banda base:")
	for {
		for i, mod := range BasebandModulation {
			fmt.Printf("%d. %s\n", i+1, mod)
		}
		fmt.Scan(&baseband)
		if baseband == 1 || baseband == 2 || baseband == 3 {
			break
		}
		fmt.Println("Método de modulação inválido. Por favor, selecione um dos métodos abaixo (1 | 2 | 3).")
	}
	var V float64
	fmt.Println("Digite o valor de modulação:")
	fmt.Scan(&V)
	// Chama as funções de modulação em banda base
	var basebanded []float64
	switch baseband {
	case 1: // NRZ Polar
		basebandFlag = "NRZ Polar"
		fmt.Println("Aplicando NRZ Polar...")
		basebanded = baseband_modulation.PolarNRZModulation(data, V)
		fmt.Println("Basebanded Data:", basebanded)

	case 2: // Bipolar
		basebandFlag = "Bipolar"
		fmt.Println("Aplicando Bipolar...")
		basebanded = baseband_modulation.BipolarModulation(data, V)
		fmt.Println("Basebanded Data:", basebanded)

	case 3: // Manchester
		basebandFlag = "Manchester"
		fmt.Println("Aplicando Manchester...")
		basebanded = baseband_modulation.ManchesterModulation(data, V)
		fmt.Println("Basebanded Data:", basebanded)

	default:
		fmt.Println("Método de modulação inválido!")
	}

	// Envia a mensagem para o servidor via socket
	// Invoca socket_client.go para comunicação com o servidor
	Send(basebanded, framingFlag, carrierFlag, basebandFlag)
}
