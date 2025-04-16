package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/adzi007/ecommerce-notification-service/internal/domain"
	"github.com/adzi007/ecommerce-notification-service/internal/dto"
)

func (r *RabbitMQ) ConsumeOrderStatus(queueName string, notifUsecase domain.NotificationUsecase) error {

	err := DeclareQueue(r.Channel, queueName)
	if err != nil {
		return err
	}

	msgs, err := r.Channel.Consume(
		queueName,
		"order_consumer",
		true,  // Auto-acknowledge
		false, // Not exclusive
		false, // No local
		false, // No-wait
		nil,   // Args
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			// defer cancel()

			var orderMessage OrderMessage
			if err := json.Unmarshal(msg.Body, &orderMessage); err != nil {
				log.Printf("Failed to parse message: %v", err)
				continue
			}

			log.Printf("Received Order: %+v", orderMessage)

			n := strconv.Itoa(int(orderMessage.OrderID))

			fmt.Println("orderMessage >>>>> ", orderMessage)

			notifMessagae := &dto.NotificationData{
				UserID: orderMessage.UserId,
				Title:  "Your Order Status is " + orderMessage.Status,
				Body:   "lorem ipsum dolor sit amet",
				Link:   "http://www.example.com/order/" + n,
				Status: 1,
			}

			// notifMessagae := &dto.NotificationData{
			// 	UserID: "d86g8d7fgdf",
			// 	Title:  "Your Order Confirmed",
			// 	Body:   "lorem ipsum dolor sit amet",
			// 	Link:   "http://www.example.com/order/" + n,
			// 	Status: 1,
			// }

			notifUsecase.Insert(notifMessagae)
		}
	}()

	return nil
}
