package cmd

import (
	"github.com/peterh/liner"
)

// Lista de comandos disponibles
var comandosDisponibles = []string{
	"exit",
	"history",
	"mth",
	"cd",
	"ls",
	"pwd",
	"clear",
}

// ConfigurarAutocompletado configura la funci�n de autocompletado
func ConfigurarAutocompletado(line *liner.State) {
	line.SetCompleter(func(linea string) []string {
		var sugerencias []string
		for _, comando := range comandosDisponibles {
			if len(linea) == 0 || (len(linea) > 0 && comando[:len(linea)] == linea) {
				sugerencias = append(sugerencias, comando)
			}
		}
		return sugerencias
	})
}
