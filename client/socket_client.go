// Manages socket communication.
// Sends the encoded message to the server and optionally waits for an acknowledgment.
package client

import (
	"IP_sim/common"
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

func Send(data interface{}) {
	// Connect to server
	conn, err := net.Dial(common.Type, common.Host+":"+common.Port)
	if err != nil {
		fmt.Println("Error ao conectar ao server:", err)
		return
	}
	defer conn.Close()

	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		fmt.Println("Error sending data:", err)
	}
	fmt.Println("Mensagem enviada com sucesso.")
}

/*
// Test the client connection (for local testing)
func main() {
	// Example signal (replace with actual modulated signal)
	signal := []float64{0.1, 0.5, -0.3, 1.0, -0.8}

	fmt.Println("Sending signal:", signal)

	err := Send(signal)
	if err != nil {
		fmt.Printf("Error sending signal: %v\n", err)
	} else {
		fmt.Println("Signal sent successfully.")
	}
}
*/

/*
import (
	"IP_sim/common"
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func StartClient(wg *sync.WaitGroup, message string) {
	fmt.Println("Starting " + common.Type + " client on " + common.Host + ":" + common.Port)

	defer wg.Done()

	// Connect to server
	conn, err := net.Dial(common.Type, common.Host+":"+common.Port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Send message to server
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	// Receive response from server
	buffer := make([]byte, 1024)
	data, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Println("Response from server: ", string(buffer[:data]))

	// Envio contínuo de mensagens
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Digite uma mensagem (ou 'exit' para sair): ")
		if !scanner.Scan() {
			break
		}
		msg := scanner.Text()
		if msg == "exit" {
			fmt.Println("Encerrando conexão...")
			break
		}

		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("Erro ao enviar mensagem:", err)
			break
		}

		data, err = conn.Read(buffer)
		if err != nil {
			fmt.Println("Erro ao ler resposta do servidor:", err)
			break
		}
		fmt.Println("Resposta do servidor:", string(buffer[:data]))
	}
}
*/
