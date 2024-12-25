package main

import (
	"myproject/internal/database"
	"myproject/internal/handlers"
	"myproject/internal/taskService"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Вызываем метод InitDB() из файла db.go
	database.InitDB()

	// Автоматическая миграция модели Message
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/patch/{id}", handler.PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
