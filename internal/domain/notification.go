package domain

import (
	"context"

	"github.com/adzi007/ecommerce-notification-service/internal/dto"
)

type Notification struct {
	ID        int64  `gorm:"primaryKey"`
	UserID    string `gorm:"not null"`
	Title     string `gorm:"not null"`
	Body      string `gorm:"not null"`
	Status    int8   `gorm:"not null"`
	IsRead    int8   `gorm:"not null"`
	CreatedAt string
}

type NotificationRepository interface {
	FindByUser(userId string) ([]Notification, error)
	Insert(notification dto.NotificationData) error
	Update(notification Notification) error
}

type NotificationService interface {
	FindByUser(ctx context.Context, user string) ([]dto.NotificationData, error)
}

type NotifWebsocket interface {
	Run()
	// Join(*websocket.Conn, string)
	// Leave(*websocket.Conn)
	// Broadcast(ChatBubble, string)
	// HandleWsChatRoom() func(*websocket.Conn)
}
