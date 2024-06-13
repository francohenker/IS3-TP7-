package routes

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("routes/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
