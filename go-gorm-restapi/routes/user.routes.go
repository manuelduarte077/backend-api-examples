package routes

import (
	"encoding/json"
	"github/manuelduarte077/go-gorm-restapi/db"
	"github/manuelduarte077/go-gorm-restapi/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User // Slice de usuarios
	db.DB.Find(&users)      // SELECT * FROM users

	json.NewEncoder(w).Encode(users) // Convertimos el slice a JSON
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)            // Obtenemos los parámetros de la URL
	var user models.User             // Creamos una instancia de User
	db.DB.First(&user, params["id"]) // SELECT * FROM users WHERE id = params["id"]

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // Si no existe el usuario, devolvemos un 404
		w.Write([]byte("User not found"))  // Y un mensaje de error
		return                             // Cortamos la ejecución
	}

	db.DB.Model(&user).Association("Groups").Find(&user.Groups) // SELECT * FROM groups WHERE user_id = user.ID
	json.NewEncoder(w).Encode(user)                             // Si no hay error, devolvemos el usuario
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User                  // Creamos una instancia de User
	json.NewDecoder(r.Body).Decode(&user) // Convertimos el JSON a un objeto User
	createdUser := db.DB.Create(&user)    // Insertamos el usuario en la base de datos
	err := createdUser.Error              // Obtenemos el error de la operación

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Si hay un error, devolvemos un 400
		w.Write([]byte(err.Error()))         // Y el mensaje de error
	}

	json.NewEncoder(w).Encode(user) // Si no hay error, devolvemos el usuario creado

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)            // Obtenemos los parámetros de la URL
	var user models.User             // Creamos una instancia de User
	db.DB.First(&user, params["id"]) // SELECT * FROM users WHERE id = params["id"]

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // Si no existe el usuario, devolvemos un 404
		w.Write([]byte("User not found"))  // Y un mensaje de error
		return                             // Cortamos la ejecución
	}

	db.DB.Unscoped().Delete(&user) // DELETE FROM users WHERE id = params["id"]
	w.WriteHeader(http.StatusOK)   // Si no hay error, devolvemos un 200
}
