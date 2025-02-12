package main

import (
	"fmt"
	"os"

	"cli/cmd" // Importamos el paquete "cmd"

	"github.com/fatih/color"
	"github.com/peterh/liner"
)

// Mostrar el prompt con el estado del �ltimo comando
func MostrarPrompt() {
	var status string
	if cmd.EstadoComando { // Aseg�rate de que `EstadoComando` est� exportado en `cmd`
		status = color.GreenString("\u2714") // Comando exitoso
	} else {
		status = color.RedString("\u2716") // Comando fallido
	}

	// Obtener el directorio actual
	dir, err := os.Getwd()
	if err != nil {
		color.Red("Error obteniendo el directorio")
		return
	}

	// Mostrar el prompt con el estado del �ltimo comando
	fmt.Printf("%s %s > ", status, color.MagentaString(dir))
}

// Funci�n que inicia la shell interactiva
func salamandra() {
	color.Cyan("Bienvenido a la shell Go. Escribe 'exit' para salir.")

	line := liner.NewLiner()
	defer line.Close()

	// Cargar autocompletado
	cmd.ConfigurarAutocompletado(line)

	for {
		// Capturar el directorio actual
		dir, err := os.Getwd()
		if err != nil {
			color.Red("Error obteniendo el directorio")
			return
		}

		// Determinar el estado del �ltimo comando
		var status string
		if cmd.EstadoComando {
			status = color.GreenString("\u2714") // Comando exitoso
		} else {
			status = color.RedString("\u2716") // Comando fallido
		}

		// Imprimir el prompt manualmente con colores
		fmt.Printf("%s %s > ", status, color.MagentaString(dir))

		// Leer la entrada del usuario con soporte para autocompletado
		comando, err := line.Prompt("")
		if err != nil {
			color.Red("Error al leer la entrada: %v", err)
			return
		}

		// Guardar en historial
		line.AppendHistory(comando)

		// Ejecutar el comando usando el paquete "cmd"
		cmd.EjecutarComando(comando)
	}
}

func main() {
	salamandra()
}
