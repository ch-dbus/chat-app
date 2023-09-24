package websocket

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"chat-app/internal/model"
	"github.com/gorilla/websocket"
)

func TestHandleConnections(t *testing.T) {
	// HTTP-Testserver starten
	server := httptest.NewServer(http.HandlerFunc(HandleConnections))
	defer server.Close()

	// Starten der BroadcastMessages Goroutine
	go BroadcastMessages()

	// WebSocket-Clients erstellen
	wsURL := "ws" + server.URL[4:] // 'http' zu 'ws' ändern
	ws1, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("could not open a ws connection on test server %v", err)
	}
	defer ws1.Close()

	ws2, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("could not open a ws connection on test server %v", err)
	}
	defer ws2.Close()

	// Testnachricht senden
	testMessage := model.Message{Nickname: "Test", Text: "Hello, world!"}
	if err := ws1.WriteJSON(testMessage); err != nil {
		t.Fatalf("could not send message over ws connection %v", err)
	}

	// Testnachricht auf dem zweiten WebSocket empfangen
	receivedMessage := &model.Message{}
	if err := ws2.ReadJSON(receivedMessage); err != nil {
		t.Fatalf("could not read message from ws connection %v", err)
	}

	// Überprüfen, ob die empfangene Nachricht der gesendeten Nachricht entspricht
	if receivedMessage.Nickname != testMessage.Nickname || receivedMessage.Text != testMessage.Text {
		t.Fatalf("received message did not match sent message")
	}
}
