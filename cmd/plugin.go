package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
)

// Mapa de plugins cargados din�micamente
var plugins = make(map[string]func([]string) string)

// Cargar todos los plugins de la carpeta "./plugins"
func CargarPlugins() {
	pluginDir := "./plugins"
	files, err := os.ReadDir(pluginDir)
	if err != nil {
		fmt.Println("No se pudo leer el directorio de plugins:", err)
		return
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".so" {
			pluginPath := filepath.Join(pluginDir, file.Name())
			err := CargarPlugin(pluginPath)
			if err != nil {
				fmt.Println("Error cargando plugin:", err)
			}
		}
	}
}

// Cargar un solo plugin
func CargarPlugin(path string) error {
	// Intentar abrir el archivo del plugin
	p, err := plugin.Open(path)
	if err != nil {
		return fmt.Errorf("error al abrir el plugin: %v", err)
	}

	// Buscar la funci�n "Run" en el plugin
	runFunc, err := p.Lookup("Run")
	if err != nil {
		return fmt.Errorf("el plugin no tiene la funci�n Run: %v", err)
	}

	// Convertir la funci�n encontrada a una funci�n compatible
	run, ok := runFunc.(func([]string) string)
	if !ok {
		return fmt.Errorf("la funci�n Run tiene una firma incorrecta")
	}

	// Registrar el plugin en el mapa
	pluginName := filepath.Base(path) // Usa el nombre del archivo como clave
	plugins[pluginName] = run
	fmt.Println("Plugin cargado:", pluginName)
	return nil
}

// Ejecutar un plugin si existe
func EjecutarPlugin(name string, args []string) {
	if run, exists := plugins[name]; exists {
		result := run(args)
		fmt.Println(result)
	} else {
		fmt.Println("Plugin no encontrado:", name)
	}
}
