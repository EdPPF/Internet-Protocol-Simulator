// Handles the main workflow -
// Captures user input (text and GUI-selected options).
// Invokes codification and modulation methods based on user selections.
// Transmits the modulated signal via socket_client.go.
package client

import (
	"IP_sim/common/communication"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Run() {
	fmt.Println("Iniciando cliente...")
	// Mostra protocolos disponíveis
	fmt.Println("Protocolos disponíveis:")
	for protocolName := range communication.SupportedProtocols {
		fmt.Println("-", protocolName)
	}

	reader := bufio.NewReader(os.Stdin)

	// User input: Seleciona protocolo de codificação
	fmt.Print("Selecione o protocolo de enquadramento: ")
	framingProtocolName, _ := reader.ReadString('\n')
	framingProtocolName = strings.TrimSpace(framingProtocolName)

	framingProtocol, ok := communication.SupportedProtocols[framingProtocolName]
	if !ok {
		fmt.Println("Protocolo de enquadramento inválido.")
		return
	}

	// User input: Seleciona protocolo de modulação (banda base ou portadora)
	fmt.Print("Selecione o protocolo de modulação: ")
	modulationProtocolName, _ := reader.ReadString('\n')
	modulationProtocolName = strings.TrimSpace(modulationProtocolName)

	modulationProtocol, ok := communication.SupportedProtocols[modulationProtocolName]
	if !ok {
		fmt.Println("Protocolo de modulação inválido.")
		return
	}

	// User input: Digita a mensagem a ser transmitida
	fmt.Print("Digite a mensagem a ser transmitida: ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)

	// Codifica a mensagem - ENQUADRAMENTO
	encodedMessage, err := framingProtocol.Apply([]byte(message))
	if err != nil {
		fmt.Println("Erro na codificação da mensagem:", err)
		return
	}

	// Modula a mensagem
	modulatedMessage, err := modulationProtocol.Apply(encodedMessage)
	if err != nil {
		fmt.Println("Erro na modulação da mensagem:", err)
		return
	}

	// Envia a mensagem modulada via socket
	fmt.Println("Enviando mensagem modulada...")
	Send(modulatedMessage)
}

func GetUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}

func GetSelectedCodification() communication.Protocol {
	input := GetUserInput()
	return communication.SupportedProtocols[input]
}

func GetSelectedModulation() communication.Protocol {
	return communication.SupportedProtocols["Manchester"]
}
