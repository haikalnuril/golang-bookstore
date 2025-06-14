package rabbitmq

import (
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Publish(queueName string, body []byte) error {
	q, err := Channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	return Channel.PublishWithContext(
		nil,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
			Timestamp:    time.Now(),
		},
	)
}
