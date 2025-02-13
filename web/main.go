package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Aseg�rate de que las rutas sin especificar un archivo dirijan al archivo HTML
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "index.html")
	} else {
		// Si el archivo solicitado est� en la carpeta est�tica, lo servimos directamente
		http.ServeFile(w, r, "."+r.URL.Path)
	}
}

func main() {
	// Ruta para servir los archivos est�ticos (como im�genes, CSS, JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handler)
	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
