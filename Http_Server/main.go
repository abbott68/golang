package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleAPIRequest)

	log.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handleAPIRequest(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string `json:"message"`
	}{
		Message: "Hello from Golang backend!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
