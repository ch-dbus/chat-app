// /chat-app/main.go
package main

import (
	"net/http"
	"chat-app/internal/handler"
	"chat-app/internal/repository"
	"chat-app/pkg/websocket"
)

func main() {
	// Datenbank initialisieren
	if err := repository.InitializeDB(); err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Hintergrundroutine zum Senden von Broadcast-Nachrichten
	go websocket.BroadcastMessages()

	// HTTP-Handler für die verschiedenen Routen festlegen
	http.HandleFunc("/message", handler.SendMessage)
	http.HandleFunc("/messages", handler.GetMessages)
	http.HandleFunc("/ws", websocket.HandleConnections)  // WebSocket-Handler

	// Starten des Servers
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
