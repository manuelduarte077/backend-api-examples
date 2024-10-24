package main

import (
	"go-api/controllers"
	"go-api/database"
	"go-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Inicializar la base de datos
	database.InitDB()

	// Crear enrutador
	router := mux.NewRouter()

	// Rutas para el CRUD de usuarios
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	// Middleware de logging
	router.Use(middleware.LoggingMiddleware)

	// Iniciar servidor
	log.Println("Servidor corriendo en el puerto 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
