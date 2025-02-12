package cmd

import (
	"bufio"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

// EjecutarComando ejecuta un comando ingresado por el usuario en la shell.
func EjecutarComando(comando string) {
	// Si el usuario escribe "exit", la shell termina.
	if strings.TrimSpace(comando) == "exit" {
		color.Yellow("Saliendo de la shell...")
		os.Exit(0)
	}

	// Crear el comando para ejecutarlo en una shell
	cmd := exec.Command("sh", "-c", comando)

	// Capturar la salida est�ndar del comando
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		color.Red("Error al capturar la salida: %v", err)
		return
	}

	// Ejecutar el comando
	err = cmd.Start()
	if err != nil {
		color.Red("Error al iniciar el comando: %v", err)
		return
	}

	// Leer la salida y mostrarla en color verde
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		color.Green(scanner.Text())
	}

	// Esperar a que el comando termine
	err = cmd.Wait()
	if err != nil {
		color.Red("Error al esperar el comando: %v", err)
	}
}
