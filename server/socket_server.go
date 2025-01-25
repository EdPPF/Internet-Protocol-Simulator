package server

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Payload struct {
	ModulatedSignal []float64
	FramingFlag     string
	CarrierFlag     string
	BasebandFlag    string
	Amplitude       float64 // Para QAM8 e ASK
	Frequency       float64 // Para QAM8 e ASK
	F1              float64 // Low Freq Para FSK
	F2              float64 // High Freq Para FSK
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Recebe a mensagem do cliente
	decoder := gob.NewDecoder(conn)
	var payload Payload
	if err := decoder.Decode(&payload); err != nil {
		fmt.Println("Erro ao receber a mensagem:", err)
		return
	}
	fmt.Printf("Mensagem recebida: %v\n", payload.ModulatedSignal)

	// Processa o payload
	originalMessage, err := ProcessPayload(payload)
	if err != nil {
		fmt.Println("Erro ao processar o payload:", err)
		return
	}

	// Exibe a mensagem original
	fmt.Printf("Mensagem original: %s\n", originalMessage)

	// Envia um ack para o cliente
	ack := "Mensagem recebida."
	encoder := gob.NewEncoder(conn)
	if err := encoder.Encode(ack); err != nil {
		fmt.Println("Erro ao enviar o ack:", err)
		return
	}
	fmt.Println("Ack enviado.")
}
