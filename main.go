package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var body Message
	json.NewDecoder(r.Body).Decode(&body)
	// Сохраняем задачу в базу данных
	DB.Create(&body)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(body)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Message
	DB.Find(&tasks)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", GetHandler).Methods("GET")
	router.HandleFunc("/api/hello", PostHandler).Methods("POST")

	http.ListenAndServe(":8080", router)
}
