package rabbitmq

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func DeclareQueue(ch *amqp091.Channel, queueName string) error {
	_, err := ch.QueueDeclare(
		queueName, // Queue name
		true,      // Durable
		false,     // Auto-delete
		false,     // Exclusive
		false,     // No-wait
		amqp091.Table{
			"x-queue-type": "quorum",
		}, // arguments
	)
	if err != nil {
		return err
	}
	log.Printf("Queue [%s] declared successfully", queueName)
	return nil
}
