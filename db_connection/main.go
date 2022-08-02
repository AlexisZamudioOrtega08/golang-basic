package main

import (
	"net/http"

	"github.com/AlexisZamudioOrtega08/db_connection/db"
	"github.com/AlexisZamudioOrtega08/db_connection/models"
	"github.com/AlexisZamudioOrtega08/db_connection/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.User{}, models.Task{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	r.HandleFunc("/tasks", routes.PostTaskHandler).Methods("POST")
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/tasks/users/{user_id}", routes.GetTasksUserHandler).Methods("GET")
	r.HandleFunc("/tasks/users/{user_id}", routes.DeleteUsersTasksHandler).Methods("DELETE")
	http.ListenAndServe(":3000", r)
}
