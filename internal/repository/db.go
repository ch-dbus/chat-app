// /chat-app/repository/db.go
package repository

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type ConnConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

var DB *gorm.DB

func InitializeDB() error {
	connConfig := ConnConfig{
		Username: "SA",
		Password: "qbgglik535PBASEFJ!qQ",
		Host:     "192.168.0.126",
		Port:     "1433",
		Database: "chat",
	}

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		connConfig.Username,
		connConfig.Password,
		connConfig.Host,
		connConfig.Port,
		connConfig.Database,
	)
	var err error
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	return err
}
