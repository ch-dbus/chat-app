// Verbindung zum WebSocket herstellen
const ws = new WebSocket("ws://localhost:8080/ws");

// Wenn die Verbindung hergestellt ist
ws.addEventListener("open", function(event) {
  console.log("Connected to the WebSocket");
});

// Wenn eine Nachricht vom Server empfangen wird
ws.addEventListener("message", function(event) {
  const message = JSON.parse(event.data);
  displayMessage(message);
});

// Funktion zum Anzeigen einer Nachricht im Chat-Fenster
function displayMessage(message) {
  const messageList = document.getElementById("message-list");
  const messageElement = document.createElement("div");
  messageElement.textContent = message.nickname + ": " + message.text;
  messageList.appendChild(messageElement);
}

// Event-Handler für den Senden-Button
document.getElementById("send-button").addEventListener("click", function() {
  const nickname = document.getElementById("nickname").value;
  const message = document.getElementById("message").value;
  
  // Nachricht an den Server senden
  ws.send(JSON.stringify({ nickname: nickname, text: message }));

  // Textfeld für die Nachricht leeren
  document.getElementById("message").value = "";
});
