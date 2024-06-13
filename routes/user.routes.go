package routes

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/francohenker/goApi/db"
	"github.com/francohenker/goApi/models"
)

// devuelve todos los usuarios de la db
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

// devuelve un usuario por el parametro ID de la pagina
func UserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	url := r.URL.String()
	regex := regexp.MustCompile("[0-9]+")

	db.DB.First(&user, regex.FindString(url))

	//esto es una parte alternativa de manejar el error cuando no se encuetra el usuario

	// err := db.DB.First(&user, regex.FindString(url)).Error
	// if err == gorm.ErrRecordNotFound {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	//si no se encuetra un usuario devuelve ID=0, nose que tan seguro sea manejarlo asi de todas formas
	if user.ID == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&user)
}

// crea un usuario en la bd
func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)

	//se controla que los campos ingresados sean validos
	if newUser.Firstname == "" || newUser.Lastname == "" || newUser.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(error.Error(http.ErrBodyNotAllowed)))
		return
	}

	// se crea el usuario en la bd y se controlan los errores
	createNewUser := db.DB.Create(&newUser)

	err := createNewUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	//se devuelve el usuario creado en la bd
	json.NewEncoder(w).Encode(&newUser)
}

// elimina un usuario de la bd
func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Delete"))
	var oldUser models.User
	json.NewDecoder(r.Body).Decode(&oldUser)

	// deleteUser := db.DB.Delete(&oldUser)
	// fmt.Println(db.DB.Find(&oldUser))

	err := db.DB.Where("firstname = ? AND lastname = ? AND email = ?", oldUser.Firstname, oldUser.Lastname, oldUser.Email).Delete(&oldUser).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// obtiene todas las tareas de un usuario
func GetTasksUserHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	url := r.URL.String()
	regex := regexp.MustCompile("[0-9]+")

	// implementar primero la busqueda para validar que exista un usuario con el id dado por la URL

	db.DB.Where("user_id = ?", regex.FindString(url)).Find(&tasks)

	json.NewEncoder(w).Encode(&tasks)

}
