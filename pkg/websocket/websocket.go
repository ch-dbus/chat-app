// /chat-app/pkg/websocket/websocket.go
package websocket

import (
	//"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"chat-app/internal/model"
	"chat-app/internal/repository"
	"fmt"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan model.Message)    // broadcast channel

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Hier können Sie eine Überprüfung hinzufügen
	},
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		var msg model.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			fmt.Printf("Received message: %+v\n", msg)
			delete(clients, conn)
			break
		}

		

		repository.SaveMessage(msg) // Speichern der Nachricht in der Datenbank

		broadcast <- msg
	}
}

func BroadcastMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Websocket error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
