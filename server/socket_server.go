package server

import (
	"IP_sim/common"
	"encoding/gob"
	"fmt"
	"net"
)

// Recieve recebe mensage do Cliente e envia um ack de confirmação
func Recieve() {
	ln, err := net.Listen(common.Type, common.Host+":"+common.Port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Error accepting:", err.Error())
		return
	}
	defer conn.Close()

	var message []float64
	decoder := gob.NewDecoder(conn)
	err = decoder.Decode(&message)
	if err != nil {
		fmt.Println("Error decoding:", err.Error())
		return
	}

	fmt.Println("Received message:", message)

	ack := "Message received"
	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(ack)
	if err != nil {
		fmt.Println("Error encoding ack:", err.Error())
		return
	}
}
