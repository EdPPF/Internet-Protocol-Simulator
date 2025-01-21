// Receives the modulated signal from the client.
// Demodulates the signal to extract binary data.
// Decodes the binary data into the original message.
// Displays the decoded message via the GUI.
package server

import (
	"IP_sim/common"
	"IP_sim/common/communication"
	"fmt"
)

func Run() {
	fmt.Println("Iniciando servidor...")
	fmt.Println("Escutando na porta", common.Port, "...")
	StartServer(func(data interface{}) {
		fmt.Println("Recebido:", data)

		// User input: Seleciona protocolo de demodulação
		var demodulationProtocolName string
		fmt.Print("Selecione o protocolo de demodulação:")
		fmt.Scanln(&demodulationProtocolName)

		demodulationProtocol, ok := communication.SupportedProtocols[demodulationProtocolName]
		if !ok {
			fmt.Println("Protocolo de demodulação inválido.")
			return
		}

		// Demodula a mensagem
		demodulatedMessage, err := demodulationProtocol.Apply(data)
		if err != nil {
			fmt.Println("Erro na demodulação da mensagem:", err)
			return
		}

		// User input: Seleciona protocolo de decodificação
		var decodingProtocolName string
		fmt.Print("Selecione o protocolo de decodificação: ")
		fmt.Scanln(&decodingProtocolName)

		decodingProtocol, ok := communication.SupportedProtocols[decodingProtocolName]
		if !ok {
			fmt.Println("Protocolo de decodificação inválido.")
			return
		}

		// Decodifica a mensagem
		decodedMessage, err := decodingProtocol.Reverse(demodulatedMessage)
		if err != nil {
			fmt.Println("Erro na decodificação da mensagem:", err)
			return
		}

		fmt.Println("Mensagem recebida:", string(decodedMessage.([]byte)))
	})
}
