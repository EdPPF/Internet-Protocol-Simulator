package main

import (
	"IP_sim/client"
	"IP_sim/server"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	clearScreen()
	fmt.Println("BEM VINDO")
	fmt.Println("---------------------------------------")

	var choice int
	fmt.Println("Selecione um modo:")
	fmt.Println("1. Cliente")
	fmt.Println("2. Servidor")
	fmt.Println("3. Sair")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Escolha inválida. Por favor, selecione um número entre 1, 2 e 3.")
		return
	}

	switch choice {
	case 1: // Inicia o cliente
		fmt.Println("Iniciando cliente...")
		client.Run()

	case 2: // Inicia o servidor
		fmt.Println("Iniciando servidor...")
		server.Run()

	case 3: // Sair
		fmt.Println("Encerrando...")
		return

	default:
		fmt.Println("Escolha inválida. Por favor, selecione um número entre 1, 2 e 3.")
	}
}
