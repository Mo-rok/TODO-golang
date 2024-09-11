package utils

import (
	"fmt"
	"lessons/models"
	"net/http"
	"time"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	taskName := r.FormValue("name")
	taskDescription := r.FormValue("description")
	taskStartAtStr := r.FormValue("startAt") // Строковое значение
	taskEndAtStr := r.FormValue("finishAt")  // Строковое значение
	taskToListName := r.FormValue("listName")

	// Парсинг времени из строки
	taskStartAt, err := time.Parse(time.RFC3339, taskStartAtStr)
	if err != nil {
		http.Error(w, "Invalid start date format", http.StatusBadRequest)
		return
	}

	taskEndAt, err := time.Parse(time.RFC3339, taskEndAtStr)
	if err != nil {
		http.Error(w, "Invalid end date format", http.StatusBadRequest)
		return
	}

	// Создание списка (в реальном приложении это скорее всего будет извлечение из базы данных)
	taskList := &models.List{Name: taskToListName}

	newTask := CreateTask(taskName, taskDescription, taskStartAt, taskEndAt, taskList)

	fmt.Fprintf(w, "Has been created new task: %+v", newTask)
}

func CreateTask(name string, description string, startAt time.Time, endAt time.Time, list *models.List) models.Task {
	return models.Task{
		Name:        name,
		Description: description,
		StartAt:     startAt,
		EndAt:       endAt,
		List:        list,
	}
}
