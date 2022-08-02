package routes

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/AlexisZamudioOrtega08/db_connection/db"
	"github.com/AlexisZamudioOrtega08/db_connection/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []models.User{}
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user := models.User{}
	db.DB.First(&user, id)
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"msg": "User not found"})
		return
	}
	json.NewEncoder(w).Encode(user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	if len(user.FirstName) == 0 || len(user.LastName) == 0 || len(user.Email) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"msg": "Missing required fields"})
		return
	}
	createdUser := db.DB.Create(&user)
	if createdUser.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		if strings.Contains(createdUser.Error.Error(), "duplicate") {
			json.NewEncoder(w).Encode(map[string]string{"msg": "User already exists"})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"msg": "Error creating user"})
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user := models.User{}
	db.DB.First(&user, id)
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"msg": "User not found"})
		return
	}
	db.DB.Delete(&user)
	user = models.User{}
	db.DB.First(&user, id)
	if user.ID != 0 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"msg": "Error deleting user"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"msg": "User deleted"})
}
