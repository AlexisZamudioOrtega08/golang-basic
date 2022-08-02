package routes

import (
	"encoding/json"
	"net/http"

	"github.com/AlexisZamudioOrtega08/db_connection/db"
	"github.com/AlexisZamudioOrtega08/db_connection/models"
	"github.com/gorilla/mux"
)

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	task := models.Task{}
	json.NewDecoder(r.Body).Decode(&task)
	if len(task.Title) == 0 || task.UserID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"msg": "Missing required fields"})
		return
	}
	user := models.User{}
	db.DB.Where("id = ?", task.UserID).First(&user)
	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"msg": "User not found"})
		return
	}
	db.DB.Create(&task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"msg": "Task created"})
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := []models.Task{}
	db.DB.Find(&tasks)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	task := models.Task{}
	db.DB.Where("id = ?", id).First(&task)
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"msg": "Task not found"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func GetTasksUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	tasks := []models.Task{}
	db.DB.Where("user_id = ?", userID).Find(&tasks)
	if len(tasks) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"msg": "Tasks not found"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func DeleteUsersTasksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]
	db.DB.Where("user_id = ?", user_id).Delete(&models.Task{})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"msg": "Tasks deleted"})
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	task := models.Task{}
	db.DB.Where("id = ?", id).First(&task)
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"msg": "Task not found"})
		return
	}
	db.DB.Delete(&task)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"msg": "Task deleted"})
}
