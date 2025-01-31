package usecase

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/adzi007/ecommerce-notification-service/internal/domain"
	"github.com/gofiber/contrib/websocket"
)

type NotificationUsecase struct {
	// repo    *repository.NotificationRepositoryStruct
	repo    domain.NotificationRepository
	Clients map[string]*websocket.Conn
}

func NewNotificationUsecase(repo domain.NotificationRepository) *NotificationUsecase {
	return &NotificationUsecase{
		repo:    repo,
		Clients: make(map[string]*websocket.Conn),
	}
}

// WebSocket Notification
func (uc *NotificationUsecase) SendWebSocketNotification(userID string, message string) {
	client, exists := uc.Clients[userID]
	if exists {
		client.WriteMessage(websocket.TextMessage, []byte(message))
	}
}

// Email Notification
func (uc *NotificationUsecase) SendEmailNotification(email string, message string) {
	// Simulated SMTP config (use a real SMTP server)
	auth := smtp.PlainAuth("", "your-email@example.com", "your-password", "smtp.example.com")
	err := smtp.SendMail("smtp.example.com:587", auth, "your-email@example.com", []string{email}, []byte(message))
	if err != nil {
		log.Println("Failed to send email:", err)
	}
	fmt.Println("📧 Email sent to", email)
}
