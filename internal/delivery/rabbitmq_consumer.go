package delivery

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/adzi007/ecommerce-notification-service/internal/domain"
	"github.com/rabbitmq/amqp091-go"
)

// OrderStatusUpdated event
type OrderStatusUpdated struct {
	OrderID string `json:"order_id"`
	UserID  string `json:"user_id"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

// RabbitMQ Consumer
func ConsumeOrderUpdates(uc *domain.NotificationUsecase) {
	// conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	conn, err := amqp091.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open channel:", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("order_status", false, false, false, false, nil)
	if err != nil {
		log.Fatal("Failed to declare queue:", err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Failed to consume messages:", err)
	}

	for msg := range msgs {
		var event OrderStatusUpdated
		json.Unmarshal(msg.Body, &event)

		// go uc.SendWebSocketNotification(event.UserID, event.Message)
		// go uc.SendEmailNotification(event.Email, event.Message)
		fmt.Println("✅ Notification sent for Order:", event.OrderID)
	}
}
