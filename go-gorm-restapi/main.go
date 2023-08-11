package main

import (
	"github/manuelduarte077/go-gorm-restapi/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	/// Creamos el router
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)

}
