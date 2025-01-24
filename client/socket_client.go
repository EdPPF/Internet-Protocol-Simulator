package client

import (
	"IP_sim/common"
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

// Payload é a estrutura da mensagem enviada ao servidor
type Payload struct {
	ModulatedSignal []float64
	FramingFlag     string
	CarrierFlag     string
	BasebandFlag    string
}

// Semd conecta ao servidor e envia uma mensagem modulata do tipo []float64
func Send(modulatedSignal []float64, framingFlag, carrierFlag, basebandFlag string) {
	// Conecta ao servidor
	conn, err := net.Dial(common.Type, common.Host+":"+common.Port)
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}
	defer conn.Close()

	for {
		payload := Payload{
			ModulatedSignal: modulatedSignal,
			FramingFlag:     framingFlag,
			CarrierFlag:     carrierFlag,
			BasebandFlag:    basebandFlag,
		}

		// Encode the payload using gob
		var buf bytes.Buffer
		encoder := gob.NewEncoder(&buf)
		if err := encoder.Encode(payload); err != nil {
			fmt.Println("Erro ao codificar a mensagem:", err)
			return
		}

		// Envia a mensagem ao servidor
		if _, err := conn.Write(buf.Bytes()); err != nil {
			fmt.Println("Erro ao enviar a mensagem:", err)
			return
		}
		fmt.Println("Mensagem enviada com sucesso!")

		// Recebe o ack do servidor
		decoder := gob.NewDecoder(conn)
		var response string
		if err := decoder.Decode(&response); err != nil {
			fmt.Println("Erro ao receber o acknowledge do servidor:", err)
			return
		}
		fmt.Println("Ack recebido:", response)
	}
}
