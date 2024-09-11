package main

import (
	"lessons/utils"
	"log"
	"net/http"
)

func main() {
	log.Println("Server has been started...")

	// Регистрация обработчиков для маршрутов
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the homepage!"))
	})
	http.HandleFunc("/createNewList", utils.ListHandler)
	http.HandleFunc("/createNewTask", utils.TaskHandler)

	// Запуск сервера
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
