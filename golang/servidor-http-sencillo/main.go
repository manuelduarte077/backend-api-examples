package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Datos para la plantilla
type Datos struct {
	Name string
}

func main() {
	// Servir archivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Manejar la ruta raíz "/"
	http.HandleFunc("/", manejadorFormulario) // GET y POST

	// Iniciar servidor
	fmt.Println("Servidor iniciado en http://localhost:8180")
	log.Fatal(http.ListenAndServe(":8180", nil))
}

// manejadorFormulario procesa las solicitudes GET y POST
func manejadorFormulario(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Procesar datos del formulario
		name := r.FormValue("name")
		datos := Datos{Name: name}
		renderTemplate(w, "index.html", datos)
	} else {
		// Mostrar el formulario vacío en GET
		renderTemplate(w, "index.html", Datos{})
	}
}

// renderTemplate renderiza las plantillas HTML
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("templates", tmpl)
	tmplParsed, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "No se pudo cargar la plantilla", http.StatusInternalServerError)
		return
	}
	tmplParsed.Execute(w, data)
}
