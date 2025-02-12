package main

import (
	"flag"
	"fmt"
)

func main() {
	//Args:

	nombre := flag.String("nombre", "", "El nombre de usuario")
	saludar := flag.Bool("saludar", false, "Saludar al usuario")

	//Parser args:
	flag.Parse()

	//CLI logic;
	if *saludar {
		if *nombre != "" {
			fmt.Printf("¡Hola, %s!\n", *nombre)
		} else {
			fmt.Println("Hola Mundo!")
		}
	} else {
		if *nombre != "" {
			fmt.Printf("Nombre proporcionado: %s\n", *nombre)

		} else {
			fmt.Println("No se a proporcionado un nombre.")
		}
	}
}