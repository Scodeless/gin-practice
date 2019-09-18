package controller

import (
	"gin-practice/pkg/rabbitMQ"
	"github.com/streadway/amqp"
)

type RabbitMQController struct {
	BaseController
}

func (c *RabbitMQController) Producer() {
	ch, err := rabbitMQ.RMQ.Channel()
	if err != nil {
		c.FailResponse(err)
		return
	}
	defer ch.Close()
	queue, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		c.FailResponse(err)
		return
	}
	body := c.GinContext.PostForm("message")
	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		c.FailResponse(err)
		return
	}
	ch.Close()
	c.SuccessResponse(nil)
}

func (c *RabbitMQController) Consumer() {
	ch, err := rabbitMQ.RMQ.Channel()
	if err != nil {
		c.FailResponse(err)
		return
	}
	defer ch.Close()

	//declare consumer queue
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	if err != nil {
		c.FailResponse(err)
		return
	}

	//consume queue
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		c.FailResponse(err)
		return
	}
	sliceBuf := make([]string, 0, 32)
	//loop consume queue to get message
	go func() {
		for d := range msgs {
			sliceBuf = append(sliceBuf, string(d.Body))
		}
	}()

	c.SuccessResponse(sliceBuf)
}
