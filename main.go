package main

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
)

// Funci�n de saludo
func saludar(nombre string) {
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	if nombre != "" {
		fmt.Printf(green("�Hola, %s!\n"), cyan(nombre))
	} else {
		fmt.Println("�Hola, mundo!")
	}
}

// Funci�n para mostrar la ayuda personalizada
func mostrarAyuda() {
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Println(blue("Uso: mi-cli [opciones]"))
	fmt.Println("\nOpciones:")
	fmt.Println("  --nombre  <nombre>    Saludar con el nombre proporcionado")
	fmt.Println("  --saludar            Saludar al usuario")
	fmt.Println("  --help               Mostrar esta ayuda")
}

func main() {
	// Crear los par�metros de la l�nea de comandos
	nombre := flag.String("nombre", "", "El nombre del usuario")
	activarSaludo := flag.Bool("saludar", false, "Activar el saludo")
	help := flag.Bool("help", false, "Mostrar ayuda")

	// Parsear los argumentos
	flag.Parse()

	// Si el usuario solicita ayuda, mostrarla
	if *help {
		mostrarAyuda()
		return
	}

	// Si el flag --saludar es true, llamar a la funci�n saludar
	if *activarSaludo {
		saludar(*nombre)
	} else if *nombre != "" {
		// Si solo se pasa el nombre, mostrarlo
		fmt.Printf("Nombre proporcionado: %s\n", *nombre)
	} else {
		// Si no se pasa ning�n par�metro, mostrar un mensaje predeterminado
		color.Red("No se ha proporcionado un nombre ni el comando --saludar.")
	}
}
