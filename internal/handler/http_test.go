package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func TestSendMessage(t *testing.T) {
	req, err := http.NewRequest("POST", "/message", strings.NewReader(`{"nickname":"test", "text":"hello"}`))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}


func TestGetMessages(t *testing.T) {
	// Erstellen eines neuen HTTP-GET-Request
	req, err := http.NewRequest("GET", "/messages", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Erstellen eines ResponseRecorders, um die Antwort zu speichern
	rr := httptest.NewRecorder()

	// Erstellen eines Handlers für den Test
	handler := http.HandlerFunc(GetMessages)

	// Ausführen des Handlers
	handler.ServeHTTP(rr, req)

	// Überprüfen des HTTP-Statuscodes
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Optional: Überprüfen des Antwortkörpers
	// Diese Überprüfung kann kompliziert sein, da sie davon abhängt, welche Daten in der Datenbank vorhanden sind.
	// Eine Möglichkeit ist, die Datenbank vor dem Test mit bekanntem Beispielinhalt zu füllen
	// und dann zu überprüfen, ob die Antwort diesem entspricht.
}
