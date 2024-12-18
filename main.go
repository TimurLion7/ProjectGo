package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var message Message

	DB.First(&message, id)

	DB.Delete(&message)

	w.WriteHeader(http.StatusNoContent)
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updates map[string]interface{}

	json.NewDecoder(r.Body).Decode(&updates)

	DB.Model(&Message{}).Where("id = ?", id).Updates(updates)

	var updateMessage Message

	DB.First(&updateMessage, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateMessage)
}

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
	router.HandleFunc("/api/get", GetHandler).Methods("GET")
	router.HandleFunc("/api/post", PostHandler).Methods("POST")
	router.HandleFunc("/api/patch/{id}", PatchHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", DeleteHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
