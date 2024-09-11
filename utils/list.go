package utils

import (
	"lessons/models"
	"net/http"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	listName := r.FormValue("name")
	CreateList(listName)
}

func CreateList(name string) models.List {
	return models.List{
		Name: name,
	}
}
