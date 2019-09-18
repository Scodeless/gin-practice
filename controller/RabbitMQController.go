package controller

import (
	"bytes"
	"fmt"
	"gin-practice/pkg/rabbitMQ"
	"github.com/streadway/amqp"
	"log"
	"time"
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

/**
消费队列，作为监听服务,不可作为接口访问
*/
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
	//loop consume queue to get message
	//go func() {
	for d := range msgs {
		fmt.Println(string(d.Body))
		//sliceBuf = append(sliceBuf, string(d.Body))
	}
	//}()

	c.SuccessResponse(nil)
}

/**
发送工作队列消息
*/
func (c *BaseController) WorkerQueueProducer() {
	body := c.GinContext.PostForm("message")
	ch, err := rabbitMQ.RMQ.Channel()
	if err != nil {
		c.FailResponse(err)
		return
	}

	queue, err := ch.QueueDeclare("hello-work-queue", false, false, false, false, nil)
	if err != nil {
		c.FailResponse(err)
		return
	}
	err = ch.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/explain",
		Body:         []byte(body),
	})
	if err != nil {
		c.FailResponse(err)
		return
	}
	c.SuccessResponse("sent message success")
}

func (c *RabbitMQController) WorkerQueueConsumer() {
	ch, err := rabbitMQ.RMQ.Channel()
	if err != nil {
		c.FailResponse(err)
		return
	}
	defer ch.Close()

	//declare consumer queue
	q, err := ch.QueueDeclare("hello-work-queue", false, false, false, false, nil)
	if err != nil {
		c.FailResponse(err)
		return
	}

	//consume queue
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		c.FailResponse(err)
		return
	}
	//loop consume queue to get message
	//go func() {
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		dotCount := bytes.Count(d.Body, []byte("."))
		t := time.Duration(dotCount)
		time.Sleep(t * time.Second)
		log.Printf("Done")
		_ = d.Ack(false)
	}
	//}()

	c.SuccessResponse(nil)
}
