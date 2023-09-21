package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responsdWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithErr(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 5xx error: ", message)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	responsdWithJSON(w, code, errResponse{
		Error: message,
	})
}