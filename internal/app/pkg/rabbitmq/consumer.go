package rabbitmq

import (
	"encoding/json"
	"log"
)

type OrderNotification struct {
	OrderID string `json:"order_id"`
	UserID  string `json:"user_id"`
	Total   int    `json:"total_price"`
}

func StartOrderConsumer() {
	q, err := Channel.QueueDeclare("order.notification", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Queue declare failed: %v", err)
	}

	msgs, err := Channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to consume: %v", err)
	}

	go func() {
		for d := range msgs {
			var notif OrderNotification
			if err := json.Unmarshal(d.Body, &notif); err != nil {
				log.Printf("Failed to parse: %v", err)
				continue
			}
			log.Printf("🔔 New Order: ID=%s, User=%s, Total=%d", notif.OrderID, notif.UserID, notif.Total)
		}
	}()
}
