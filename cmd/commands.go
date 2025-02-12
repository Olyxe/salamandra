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
	history = append(history, comando)

	if strings.TrimSpace(comando) == "exit" {
		color.Yellow("Saliendo de la shell...")
		os.Exit(0)
	}

	if strings.TrimSpace(comando) == "history" {
		color.Cyan("Historial de comandos:")
		for i, cmd := range history {
			color.Green("%d: %s\n", i+1, cmd)
		}
		return
	}

	// Verificar si el comando es "mth" para operaciones matem�ticas
	if strings.HasPrefix(comando, "mth") {
		EjecutarOperacionMatematica(comando)
		return
	}

	// Si no es un comando interno, ejecutar como comando externo
	cmd := exec.Command("sh", "-c", comando)

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

	err = cmd.Start()
	if err != nil {
		color.Red("Error al iniciar el comando: %v", err)
	}

	scannerOut := bufio.NewScanner(stdout)
	for scannerOut.Scan() {
		color.Green(scannerOut.Text())
	}

	scannerErr := bufio.NewScanner(stderr)
	for scannerErr.Scan() {
		color.Red(scannerErr.Text())
	}

	err = cmd.Wait()
	if err != nil {
		color.Red("Error al ejecutar el comando: %v", err)
	}
}
