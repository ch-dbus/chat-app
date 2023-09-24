package repository

import (
	"os"
	"testing"
)

func TestInitializeDB(t *testing.T) {
	// Setzen Sie Umgebungsvariablen für den Test
	os.Setenv("DB_USERNAME", "SA")
	os.Setenv("DB_PASSWORD", "qbgglik535PBASEFJ!qQ")
	os.Setenv("DB_HOST", "192.168.0.126")
	os.Setenv("DB_PORT", "1433")
	os.Setenv("DB_NAME", "chat")

	err := InitializeDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Sie könnten auch weitere Überprüfungen hinzufügen,
	// z.B. ob DB *gorm.DB nicht nil ist, um sicherzustellen, dass eine Verbindung hergestellt wurde.
	if DB == nil {
		t.Fatalf("DB should not be nil")
	}
}

func TestInitializeDB_Failure(t *testing.T) {
	// Ungültige Umgebungsvariablen setzen, um einen Fehler zu provozieren
	os.Setenv("DB_USERNAME", "invalid")

	err := InitializeDB()
	if err == nil {
		t.Fatalf("Expected an error, got none")
	}

	// Überprüfen, ob die Fehlermeldung Ihren Erwartungen entspricht
	if err.Error() != "expected error message" {
		t.Fatalf("Unexpected error message: %v", err)
	}
}
