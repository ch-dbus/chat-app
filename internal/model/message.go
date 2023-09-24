// /chat-app/model/message.go
package model

type Message struct {
	ID        int    `json:"id" gorm:"column:id"`
	Nickname  string `json:"nickname" gorm:"column:nickname"`
	Message   string `json:"message" gorm:"column:message"` // Hier wurde "Text" zu "Message" geändert
	Timestamp string `json:"timestamp" gorm:"column:timestamp"`
}
