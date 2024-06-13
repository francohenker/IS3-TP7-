package main

import (
	"net/http"

	"github.com/francohenker/goApi/db"
	"github.com/francohenker/goApi/routes"
	"github.com/gorilla/mux"
)

// GetUserHandler
func main() {
	db.DBConnection()

	//MIGRACIONES DE LA BD
	// db.DB.AutoMigrate(models.Task{})
	// db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	//routes of users
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.UserHandler).Methods("GET")
	r.HandleFunc("/users/{id}/", routes.UserHandler).Methods("GET")
	r.HandleFunc("/users/{id}/tasks", routes.GetTasksUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}/tasks/", routes.GetTasksUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users", routes.DeleteUsersHandler).Methods("DELETE")
	r.HandleFunc("/users/", routes.DeleteUsersHandler).Methods("DELETE")

	//routes of tasks
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.UpdateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}/", routes.UpdateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}/", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.DeleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/tasks/", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":8000", r)

}
