package main

import (
    "github/manuelduarte077/go-gorm-restapi/db"
    "github/manuelduarte077/go-gorm-restapi/models"

    "github/manuelduarte077/go-gorm-restapi/routes"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {

    /// Conectamos a la base de datos
    db.DBConnect()

    /// Ejectuamos las migraciones
    db.DB.AutoMigrate(models.User{}, models.Task{}, models.Group{})

    /// Creamos el router
    r := mux.NewRouter()

    /// Routes
    r.HandleFunc("/", routes.HomeHandler)

    /// User router
    r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
    r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
    r.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
    r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

    /// Group router
    r.HandleFunc("/groups", routes.GetGroupsHandler).Methods("GET")
    r.HandleFunc("/groups/{id}", routes.GetGroupHandler).Methods("GET")
    r.HandleFunc("/groups", routes.CreateGroupHandler).Methods("POST")
    r.HandleFunc("/groups/{id}", routes.DeleteGroupHandler).Methods("DELETE")

    /// Task router
    r.HandleFunc("/tasks", routes.GetsTasksHandler).Methods("GET")
    r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
    r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
    r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

    http.ListenAndServe(":3000", r)

}
