package server

// Manages socket communication.
// Receives messages from the client and passes them to protocol_manager.go.

import (
	"IP_sim/common"
	"fmt"
	"net"
	"sync"
)

func StartServer(wg *sync.WaitGroup) {
	fmt.Println("Starting" + common.Type + "server on " + common.Host + ":" + common.Port)

	defer wg.Done()

	listener, err := net.Listen(common.Type, common.Host+":"+common.Port)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening on port", common.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close() // Close connection before returning

	buffer := make([]byte, 1024) // Create a buffer to hold received data. The buffer is, essetially, a byte array

	for {
		data, err := conn.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Client disconnected.")
			} else {
				fmt.Println("Error reading:", err)
			}
			return
		}

		fmt.Println("Received data: ", string(buffer[:data]))
		// Responde ao cliente
		conn.Write([]byte("Message received."))
	}
}
