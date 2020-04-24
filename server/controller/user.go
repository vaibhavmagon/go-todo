package controller

import (
	"encoding/json"
	"net/http"
	"fmt"
	"go-todo/server/models"
	"go-todo/server/service"
	"github.com/gorilla/mux"
)

// GetAllTask get all the task route
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := service.GetAllUsers()
	json.NewEncoder(w).Encode(payload)
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	payload := service.GetOneUser(params["id"])
	json.NewEncoder(w).Encode(payload)
}

// CreateTask create task route
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user, r.Body)
	service.InsertUser(user)
	json.NewEncoder(w).Encode(user)
}

// TaskComplete update task route
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	count := service.UpdateUser(params["id"], user)
	if count > 0 {
        json.NewEncoder(w).Encode(params["id"])
	}else{
        json.NewEncoder(w).Encode("User already updated!")
	}
}

// DeleteTask delete one task route
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	count := service.DeleteUser(params["id"])
	if count > 0 {
        json.NewEncoder(w).Encode(params["id"])
	}else{
        json.NewEncoder(w).Encode("No user found!")
	}
}