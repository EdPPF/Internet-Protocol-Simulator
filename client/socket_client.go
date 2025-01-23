package client

import (
	"IP_sim/common"
	"fmt"
	"net"
)

// Semd conecta ao servidor e envia uma mensagem modulata do tipo []float64
func Send(modulatedSignal []float64) {
	// Conecta ao servidor
	conn, err := net.Dial(common.Type, common.Host+":"+common.Port)
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}
	defer conn.Close()

	// Envia a mensagem modulada
	fmt.Println("Enviando mensagem modulada para o servidor...")
	_, err = fmt.Fprintf(conn, "%v\n", modulatedSignal)
	if err != nil {
		fmt.Println("Erro ao enviar mensagem modulada para o servidor:", err)
		return
	}
}
