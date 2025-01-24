package server

import (
	"IP_sim/common"
	"encoding/gob"
	"fmt"
	"net"
)

type Payload struct {
	ModulatedSignal []float64
	FramingFlag     string
	CarrierFlag     string
	BasebandFlag    string
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

	// TODO: Implementar a lógica de demodulação

	// Envia um ack para o cliente
	ack := "Mensagem recebida."
	encoder := gob.NewEncoder(conn)
	if err := encoder.Encode(ack); err != nil {
		fmt.Println("Erro ao enviar o ack:", err)
		return
	}
	fmt.Println("Ack enviado.")
}

func Start() {
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
