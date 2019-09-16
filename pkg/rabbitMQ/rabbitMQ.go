package rabbitMQ

import (
	"github.com/streadway/amqp"
	"log"
)

var RMQ *amqp.Connection

func ConnectRabbitMQServer() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("connect rabbitMQ server failed, %v", err)
	}

	RMQ = conn
	return RMQ
}