package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Funci�n que ejecuta el comando como una shell simple
func ejecutarComando(comando string) {
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
		fmt.Println(scanner.Text())
	}

	// Esperar a que termine el comando
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error al esperar el comando:", err)
	}
}

// Funci�n que maneja el bucle de la shell
func miShell() {
	fmt.Println("Bienvenido a la shell Go. Escribe 'exit' para salir.")

	// Bucle para leer comandos
	for {
		// Mostrar un prompt
		fmt.Print(">>> ")

		// Leer la entrada del usuario
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		comando := scanner.Text()

		// Si el comando es "exit", salir de la shell
		if strings.TrimSpace(comando) == "exit" {
			fmt.Println("Saliendo de la shell...")
			break
		}

		// Ejecutar el comando
		ejecutarComando(comando)
	}
}

func main() {
	// Iniciar la shell
	miShell()
}
