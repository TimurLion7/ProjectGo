package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type requestBody struct {
	Message string `json:"message"`
}

var task string

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var body requestBody
	json.NewDecoder(r.Body).Decode(&body)
	task = body.Message
	fmt.Fprintf(w, "Message: %s updated successfully!", task)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", task)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", GetHandler).Methods("GET")
	router.HandleFunc("/api/hello", PostHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
