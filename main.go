package main

import (
	"fmt"
	"os"
	"IP_sim/common/communication"
	"sync"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Uso: go run main.go [client|server]")
        return
    }

    var wg sync.WaitGroup
    wg.Add(1)

    switch os.Args[1] {
    case "client":
	var message string
	// Se comando no terminal tiver uma mensagem
	if len(os.Args) > 2 {
		message = os.Args[2]
	} else { // Se a mensagem não for passada no comando
		fmt.Println("Digite a mensagem para o servidor:")
		fmt.Scanln(&message)
	}
        go communication.StartClient(&wg, message)
    case "server":
        go communication.StartServer(&wg)
    default:
        fmt.Println("Argumento inválido. Use 'client' ou 'server'.")
        wg.Done()
    }

    wg.Wait() // Aguarda todas as goroutines finalizarem
}
