package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html", "templates/admin.html"))

func main() {
	// Inicializa la base de datos
	InitDB()

	// Configurar rutas
	r := mux.NewRouter()

	// Ruta principal - Mostrar lista de productos
	r.HandleFunc("/", HomeHandler).Methods("GET")

	// Ruta para la vista de administración
	r.HandleFunc("/admin", AdminHandler).Methods("GET")
	r.HandleFunc("/create", CreateHandler).Methods("POST")

	// Archivos estáticos para CSS y JS
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	// Iniciar servidor
	log.Println("Servidor iniciado en :8181")
	log.Fatal(http.ListenAndServe(":8181", r))
}
