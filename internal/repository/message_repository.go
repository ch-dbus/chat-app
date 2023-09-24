// /chat-app/repository/message_repository.go
package repository

import (
	"chat-app/internal/model"
)

func SaveMessage(message model.Message) error {
	result := DB.Create(&message)
	return result.Error
}

func GetLatestMessages() ([]model.Message, error) {
	var messages []model.Message
	result := DB.Order("timestamp desc").Limit(50).Find(&messages)
	return messages, result.Error
}
