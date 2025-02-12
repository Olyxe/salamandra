package cmd

import (
	"math"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// EjecutarOperacionMatematica maneja las operaciones matem�ticas desde la l�nea de comandos
func EjecutarOperacionMatematica(comando string) {
	// Separar los elementos del comando
	parts := strings.Fields(comando)

	// Verificar si el comando es "mth" y tiene los par�metros correctos
	if len(parts) != 4 || parts[0] != "mth" {
		color.Red("Uso: mth [num1] [symbol] [num2]")
		return
	}

	// Convertir los n�meros
	num1, err1 := strconv.ParseFloat(parts[1], 64)
	num2, err2 := strconv.ParseFloat(parts[3], 64)
	if err1 != nil || err2 != nil {
		color.Red("Error: Los n�meros proporcionados no son v�lidos")
		return
	}

	// Obtener el operador
	operador := parts[2]

	// Realizar la operaci�n matem�tica
	var resultado float64
	switch operador {
	case "+":
		resultado = num1 + num2
	case "-":
		resultado = num1 - num2
	case "*":
		resultado = num1 * num2
	case "/":
		if num2 == 0 {
			color.Red("Error: Divisi�n por cero no permitida")
			return
		}
		resultado = num1 / num2
	case "^":
		resultado = math.Pow(num1, num2)
	default:
		color.Red("Error: Operador no v�lido. Usa +, -, *, / o ^")
		return
	}

	// Mostrar el resultado en color verde
	color.Green("Resultado: %.2f", resultado)
}
