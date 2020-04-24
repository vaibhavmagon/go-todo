package router

import (
	"go-todo/server/controller"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/task", controller.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", controller.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", controller.GetOneTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task/{id}", controller.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", controller.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", controller.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTask", controller.DeleteAllTask).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/user", controller.GetAllUsers).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user", controller.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/{id}", controller.GetOneUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user/{id}", controller.UpdateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/user/{id}", controller.DeleteUser).Methods("DELETE", "OPTIONS")

	return router
}
