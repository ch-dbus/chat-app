package repository

import (
	"testing"
	"chat-app/internal/model"
)

func TestSaveAndGetMessages(t *testing.T) {
	// Datenbank initialisieren
	InitializeDB()

	// Nachricht speichern
	message := model.Message{
		Nickname:  "test",
		Text:      "hello world",
		Timestamp: "2023-09-21T16:00:00Z",
	}
	if err := SaveMessage(message); err != nil {
		t.Fatalf("Failed to save message: %v", err)
	}

	// Nachrichten abrufen
	messages, err := GetLatestMessages()
	if err != nil {
		t.Fatalf("Failed to get messages: %v", err)
	}

	// Überprüfen, ob die gespeicherte Nachricht abgerufen wurde
	if len(messages) == 0 {
		t.Fatalf("No messages retrieved")
	}

	lastMessage := messages[0]
	if lastMessage.Nickname != "test" || lastMessage.Text != "hello world" {
		t.Errorf("Retrieved message did not match saved message")
	}
}
