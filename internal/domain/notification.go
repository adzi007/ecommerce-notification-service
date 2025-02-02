package domain

import (
	"context"
	"time"

	"github.com/adzi007/ecommerce-notification-service/internal/dto"
	"github.com/gofiber/contrib/websocket"
)

type Notification struct {
	ID        int64  `gorm:"primaryKey"`
	UserID    string `gorm:"not null"`
	Title     string `gorm:"not null"`
	Body      string `gorm:"not null"`
	Link      string `gorm:"not null"`
	Status    int8   `gorm:"not null"`
	IsRead    int8   `gorm:"not null"`
	CreatedAt time.Time
}

type NotifMessageRequest struct {
	// ID        int64  `json:"id"`
	UserID string `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Link   string `json:"link"`
	Status int8   `json:"status"`
	IsRead int8   `json:"isRead"`
	// CreatedAt string
}

type NotificationRepository interface {
	FindByUser(userId string) ([]Notification, error)
	Insert(notification *dto.NotificationData) (Notification, error)
	Update(notification Notification) error
}

type NotificationUsecase interface {
	// FindByUser(userId string) ([]Notification, error)
	Insert(notification *dto.NotificationData) (Notification, error)
	FindByUser(userId string) ([]Notification, error)
	// Update(notification Notification) error
}

type NotificationService interface {
	FindByUser(ctx context.Context, user string) ([]dto.NotificationData, error)
}

type NotifWebsocket interface {
	Run()
	HandleNotificationRoom() func(*websocket.Conn)
	Broadcast(data Notification)
}
