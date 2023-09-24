
### Chat-App Backend-Dokumentation

---

#### Ordnerstruktur:

- `cmd/`: Enthält den Einstiegspunkt der Anwendung (`main.go`).
- `internal/`: Enthält interne Pakete, die von der Anwendung genutzt werden.
    - `handler/`: Enthält die HTTP-Handler für die REST-API.
    - `model/`: Enthält die Datenmodelle.
    - `repository/`: Enthält die Datenbankzugriffslogik.
- `pkg/`: Enthält externe Pakete.
    - `websocket/`: Enthält den WebSocket-Handler.

---

#### Dateien:

- `main.go`: Einstiegspunkt der Anwendung. Initialisiert die Datenbank und startet den HTTP-Server.
  
---

#### `internal/handler/http.go`

- `SendMessage(w http.ResponseWriter, r *http.Request)`: Nimmt eine Nachricht im JSON-Format entgegen und speichert sie in der Datenbank.
- `GetMessages(w http.ResponseWriter, r *http.Request)`: Gibt die letzten 50 Nachrichten aus der Datenbank zurück.

---

#### `internal/model/message.go`

- `Message`: Struktur, die das Datenmodell für eine Nachricht definiert. Enthält die Felder `ID`, `Nickname`, `Message` und `Timestamp`.

---

#### `internal/repository/db.go`

- `InitializeDB()`: Initialisiert die Datenbankverbindung.
  
---

#### `internal/repository/message_repository.go`

- `SaveMessage(message model.Message)`: Speichert eine Nachricht in der Datenbank.
- `GetLatestMessages()`: Holt die letzten 50 Nachrichten aus der Datenbank.

---

#### `pkg/websocket/websocket.go`

- `HandleConnections(w http.ResponseWriter, r *http.Request)`: Managt WebSocket-Verbindungen und Nachrichten.
- `BroadcastMessages()`: Sendet empfangene Nachrichten an alle verbundenen Clients.

---

#### Tests

- Unit- und Integrationstests sind für die meisten Komponenten vorhanden. Sie befinden sich jeweils im selben Verzeichnis wie die zu testenden Dateien und haben die Endung `_test.go`.

