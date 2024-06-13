package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/francohenker/goApi/db"
	"github.com/francohenker/goApi/models"
	"gorm.io/gorm"
)

// crea una tarea y se la asigna a un usuario
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	//v que el titulo y el usuario no sea nulo
	if task.Title == "" || task.UserId == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error(http.ErrBodyNotAllowed)))
		return
	}

	//se verifica que exista el usuario
	var user models.User

	if err := db.DB.Where("id = ?", task.UserId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
	}
	//se crea la tarea
	if db.DB.Create(&task).Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// borra una tarea
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	if task.ID == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if db.DB.Where("ID = ?", task.ID).Delete(&task).Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// obtiene una tarea de un usuario
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	url := r.URL.String()
	regex := regexp.MustCompile("[0-9]+")

	db.DB.First(&task, regex.FindString(url))

	if task.ID != 0 {
		json.NewEncoder(w).Encode(&task)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)

}

// actualiza una tarea
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	if task.Title == "" && task.Description == "" && !task.Done {
		return
	}

	var oldTask models.Task

	url := r.URL.String()
	regex := regexp.MustCompile("[0-9]+")

	db.DB.First(&oldTask, regex.FindString(url))
	if oldTask.Done {
		return
	}

	oldTask.Title = task.Title
	oldTask.Description = task.Description
	oldTask.Done = task.Done

	json.NewEncoder(w).Encode(&oldTask)
	db.DB.Updates(&oldTask)
}
