// cmd/comandos.go

package cmd

import (
	"bufio"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

// Historial de comandos
var history []string

// Hacemos que la variable sea exportada para poder acceder desde otros paquetes
var EstadoComando bool

// EjecutarComando ejecuta comandos externos e internos en la shell
func EjecutarComando(comando string) {
	// Guardar el comando en el historial
	history = append(history, comando)

	// Comando "exit" para salir de la shell
	if strings.TrimSpace(comando) == "exit" {
		color.Yellow("Saliendo de la shell...")
		os.Exit(0)
	}

	// Comando interno: "history"
	if strings.TrimSpace(comando) == "history" {
		color.Cyan("Historial de comandos:")
		for i, cmd := range history {
			color.Green("%d: %s\n", i+1, cmd)
		}
		return
	}

	// Crear el comando para ejecutarlo en una shell
	cmd := exec.Command("sh", "-c", comando)

	// Capturar la salida est�ndar y de error
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		color.Red("Error al capturar la salida: %v", err)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		color.Red("Error al capturar el error: %v", err)
		return
	}

	// Ejecutar el comando
	err = cmd.Start()
	if err != nil {
		color.Red("Error al iniciar el comando: %v", err)
	}

	// Leer la salida est�ndar en verde
	scannerOut := bufio.NewScanner(stdout)
	for scannerOut.Scan() {
		color.Green(scannerOut.Text())
	}

	// Leer la salida de error en rojo
	scannerErr := bufio.NewScanner(stderr)
	for scannerErr.Scan() {
		color.Red(scannerErr.Text())
	}

	// Esperar a que el comando termine
	err = cmd.Wait()
	if err != nil {
		color.Red("Error al ejecutar el comando: %v", err)
	} else {
		// Si el comando fue exitoso, actualizamos el estado
		EstadoComando = true
	}
}
