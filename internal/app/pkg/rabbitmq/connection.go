package rabbitmq

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

func Init() {
	var err error
	Conn, err = amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	Channel, err = Conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}
}

func Close() {
	Channel.Close()
	Conn.Close()
}