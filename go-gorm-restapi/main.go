package main

import (
	"github/manuelduarte077/go-gorm-restapi/db"
	"github/manuelduarte077/go-gorm-restapi/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	/// Conectamos a la base de datos
	db.DBConnect()

	/// Creamos el router
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)

}
