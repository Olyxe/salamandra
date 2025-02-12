package main

import (
	"bufio"
	"fmt"
	"os"

	"cli/cmd" // Importamos el paquete "cmd"

	"github.com/fatih/color"
)

func MostrarPrompt() {
	var status string
	if cmd.estadoComando {
		status = color.GreenString("\u2714") // Comando exitoso
	} else {
		status = color.RedString("\u2716") // Comando fallido
	}

	// Obtener el directorio actual
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error obteniendo el directorio")
		return
	}

	// Mostrar el prompt con el estado del �ltimo comando
	fmt.Printf("%s %s > ", status, color.MagentaString(dir))
}

// Funci�n que inicia la shell interactiva
func miShell() {
	color.Cyan("Bienvenido a la shell Go. Escribe 'exit' para salir.")

	for {
		MostrarPrompt()
		// Obtener el directorio actual
		dir, err := os.Getwd()
		if err != nil {
			color.Red("Error al obtener el directorio actual: %v", err)
			return
		}

		flg := color.YellowString("+main")

		// Mostrar el prompt con el directorio actual en color magenta
		fmt.Print(color.MagentaString("%s > ", dir), flg, " ")

		// Leer la entrada del usuario
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		comando := scanner.Text()

		// Ejecutar el comando usando el paquete "cmd"
		cmd.EjecutarComando(comando)
	}
}

func main() {
	miShell()
}
