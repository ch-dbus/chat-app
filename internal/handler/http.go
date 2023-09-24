// /chat-app/handler/http.go
package handler

import (
	"encoding/json"
	"net/http"
	"chat-app/internal/model"
	"chat-app/internal/repository"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var message model.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if message.Nickname == "" || message.Message == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
		return
	}	

	if err := repository.SaveMessage(message); err != nil {
		http.Error(w, "Failed to save message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}


func GetMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := repository.GetLatestMessages()
	if err != nil {
		http.Error(w, "Failed to get messages", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(messages)
}
