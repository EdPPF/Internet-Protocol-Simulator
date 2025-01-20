package main

import (
	"IP_sim/common/communication"
	"fmt"
	"os"
	"sync"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go [client|server]")
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)

	switch os.Args[1] {
	case "client":

		window.SetWindowTitle("Cliente")

		var message string
		// Se comando no terminal tiver uma mensagem
		if len(os.Args) > 2 {
			message = os.Args[2]
		}
		label := widgets.NewQLabel(nil, 0)
		label.SetText(fmt.Sprintf("Digite a mensagem para o servidor:"))

		// input := widgets.NewQLineEdit(message)
		// input.SetPlaceholderText("Mensagem")

		window.SetCentralWidget(label)

		// Show the window
		window.Resize(core.NewQSize2(400, 300))
		window.Show()

		// Start the application
		app.Exec()
		go communication.StartClient(&wg, message)
	case "server":
		window.SetWindowTitle("Servidor")

		label := widgets.NewQLabel(nil, 0)
		window.SetCentralWidget(label)

		// Show the window
		window.Resize(core.NewQSize2(400, 300))
		window.Show()

		// Start the application
		app.Exec()

		go communication.StartServer(&wg)
	default:
		fmt.Println("Argumento inválido. Use 'client' ou 'server'.")
		wg.Done()
	}

}
