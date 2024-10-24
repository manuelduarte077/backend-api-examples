package routes

import (
	"encoding/json"
	"github/manuelduarte077/go-gorm-restapi/db"
	"github/manuelduarte077/go-gorm-restapi/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetsTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task // Slice de grupos
	db.DB.Find(&tasks)      // SELECT * FROM tasks

	json.NewEncoder(w).Encode(tasks) // Convertimos el slice a JSON
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task  // Creamos una instancia de Task
	params := mux.Vars(r) // Obtenemos los parámetros de la URL
	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // Si no existe el grupo, devolvemos un 404
		w.Write([]byte("Task not found"))  // Y un mensaje de error
		return                             // Cortamos la ejecución
	}

	db.DB.Model(&task).Association("Groups").Find(&task.GroupID) // SELECT * FROM tasks WHERE group_id = group.ID
	json.NewEncoder(w).Encode(&task)                             // Si no hay error, devolvemos el grupo creado
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task                  // Creamos una instancia de Task
	json.NewDecoder(r.Body).Decode(&task) // Convertimos el JSON a un objeto Task
	db.DB.Create(&task)                   // Insertamos el grupo en la base de datos
	err := db.DB.Error                    // Obtenemos el error de la operación

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Si hay un error, devolvemos un 400
		w.Write([]byte(err.Error()))         // Y el mensaje de error
	}

	json.NewEncoder(w).Encode(&task) // Si no hay error, devolvemos el grupo creado
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)            // Obtenemos los parámetros de la URL
	var task models.Task             // Creamos una instancia de Task
	db.DB.First(&task, params["id"]) // SELECT * FROM groups WHERE id = params["id"]

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // Si no existe el grupo, devolvemos un 404
		w.Write([]byte("Task not found"))  // Y un mensaje de error
		return                             // Cortamos la ejecución
	}

	db.DB.Unscoped().Delete(&task) // DELETE FROM groups WHERE id = params["id"]
	w.WriteHeader(http.StatusOK)   // Si no hay error, devolvemos un 200

}
