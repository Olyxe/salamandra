package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

// Funci�n que ejecuta el comando como una shell simple
func ejecutarComando(comando string) {
	// Si el comando es "exit", salir de la shell sin ejecutarlo
	if strings.TrimSpace(comando) == "exit" {
		color.Yellow("Saliendo de la shell...")
		os.Exit(0) // Salir del programa
	}

	// Crear el comando
	cmd := exec.Command("sh", "-c", comando)

	// Crear un buffer para capturar la salida del comando
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error al capturar la salida:", err)
		return
	}

	// Ejecutar el comando
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error al iniciar el comando:", err)
		return
	}

	// Leer y mostrar la salida del comando
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		// Colorear la salida de los comandos en verde
		color.Green(scanner.Text())
	}

	// Esperar a que termine el comando
	err = cmd.Wait()
	if err != nil {
		color.Red("Error al esperar el comando: %v", err)
	}
}

// Funci�n que maneja el bucle de la shell
func miShell() {
	// Colorear el mensaje de bienvenida
	color.Cyan("Bienvenido a la shell Go. Escribe 'exit' para salir.")

	// Bucle para leer comandos
	for {
		// Mostrar un prompt colorido en la misma l�nea
		fmt.Print("\r>>> ")

		// Leer la entrada del usuario
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		comando := scanner.Text()

		// Ejecutar el comando
		ejecutarComando(comando)
	}
}

func main() {
	// Iniciar la shell
	miShell()
}
