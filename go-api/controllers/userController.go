package controllers

import (
	"encoding/json"
	"go-api/database"
	"go-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

// Crear un nuevo usuario
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

// Obtener todos los usuarios
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// Obtener usuario por ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	database.DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

// Actualizar un usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	database.DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

// Eliminar un usuario
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	database.DB.Delete(&user, params["id"])
	w.WriteHeader(http.StatusNoContent)
}
