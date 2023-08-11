package routes

import (
	"encoding/json"
	"github/manuelduarte077/go-gorm-restapi/db"
	"github/manuelduarte077/go-gorm-restapi/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetGroupsHandler(w http.ResponseWriter, r *http.Request) {
	var groups []models.Group // Slice de grupos
	db.DB.Find(&groups)       // SELECT * FROM groups

	json.NewEncoder(w).Encode(groups) // Convertimos el slice a JSON
}

func GetGroupHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)             // Obtenemos los parámetros de la URL
	var group models.Group            // Creamos una instancia de Group
	db.DB.First(&group, params["id"]) // SELECT * FROM groups WHERE id = params["id"]

	if group.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // Si no existe el grupo, devolvemos un 404
		w.Write([]byte("Group not found")) // Y un mensaje de error
		return                             // Cortamos la ejecución
	}

	db.DB.Model(&group).Association("Tasks").Find(&group.Tasks) // SELECT * FROM tasks WHERE group_id = group.ID
	json.NewEncoder(w).Encode(group)                            // Si no hay error, devolvemos el grupo
}

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	var group models.Group                 // Creamos una instancia de Group
	json.NewDecoder(r.Body).Decode(&group) // Convertimos el JSON a un objeto Group
	createdGroup := db.DB.Create(&group)   // Insertamos el grupo en la base de datos
	err := createdGroup.Error              // Obtenemos el error de la operación

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Si hay un error, devolvemos un 400
		w.Write([]byte(err.Error()))         // Y el mensaje de error
	}

	json.NewEncoder(w).Encode(group) // Si no hay error, devolvemos el grupo creado

}

func DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)             // Obtenemos los parámetros de la URL
	var group models.Group            // Creamos una instancia de Group
	db.DB.First(&group, params["id"]) // SELECT * FROM groups WHERE id = params["id"]

	if group.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // Si no existe el grupo, devolvemos un 404
		w.Write([]byte("Group not found")) // Y un mensaje de error
		return                             // Cortamos la ejecución
	}

	db.DB.Unscoped().Delete(&group) // DELETE FROM groups WHERE id = params["id"]
	w.WriteHeader(http.StatusOK)    // Si no hay error, devolvemos un 200
}
