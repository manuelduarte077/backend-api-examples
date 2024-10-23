package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Inicializar la base de datos
func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	// Crear tabla si no existe
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			description TEXT
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

// HomeHandler - Mostrar lista de productos
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, description FROM items")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []map[string]string
	for rows.Next() {
		var id int
		var name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, map[string]string{
			"id":          strconv.Itoa(id),
			"name":        name,
			"description": description,
		})
	}

	// Renderizar el template con los productos
	tmpl.ExecuteTemplate(w, "index.html", items)
}

// AdminHandler - Muestra el formulario para crear productos
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "admin.html", nil)
}

// CreateHandler - Añadir nuevo producto
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	description := r.FormValue("description")
	_, err := db.Exec("INSERT INTO items (name, description) VALUES (?, ?)", name, description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
