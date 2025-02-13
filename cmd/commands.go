// cmd/comandos.go

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

// Historial de comandos
var history []string

// Estado del �ltimo comando ejecutado
var EstadoComando bool

// EjecutarComando ejecuta comandos externos e internos en la shell
func EjecutarComando(comando string) {
	history = append(history, comando)

	parts := strings.Fields(comando)
	if len(parts) == 0 {
		return
	}

	// Comprobar si es un plugin
	if _, existe := plugins[parts[0]]; existe {
		EjecutarPlugin(parts[0], parts[1:])
		return
	}

	switch strings.TrimSpace(comando) {
	case "exit":
		color.Yellow("Saliendo de la shell...")
		os.Exit(0)

	case "history":
		color.Cyan("Historial de comandos:")
		for i, cmd := range history {
			fmt.Printf("%d: %s\n", i+1, cmd)
		}
		return
	}

	// Verificar si el comando es "mth" para operaciones matem�ticas
	if strings.HasPrefix(comando, "mth") {
		EjecutarOperacionMatematica(comando)
		return
	}

	if strings.HasPrefix(comando, "neofetch") {
		mostrarNeofetch(comando)
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
		EstadoComando = false
		return
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
		EstadoComando = false
	} else {
		EstadoComando = true
	}
}
