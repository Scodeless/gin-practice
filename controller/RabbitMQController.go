package controller

import (
	"gin-practice/pkg/rabbitMQ"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQController struct {
	BaseController
}

func (c *RabbitMQController) Producer() {
	ch, err := rabbitMQ.RMQ.Channel()
	if err != nil {
		log.Fatalf("open rabbitMQ channel failed, %v", err)
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
		c.GinContext.JSON(-1, gin.H{
			"error": "declare queue failed," + err.Error(),
		})
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
		c.GinContext.JSON(-1, gin.H{
			"error": "publish message failed," + err.Error(),
		})
		return
	}

	c.GinContext.JSON(200, gin.H{
		"message": "publish message success",
	})
}

func (c *RabbitMQController) Consumer() {
	ch, err := rabbitMQ.RMQ.Channel()
	if err != nil {
		log.Fatalf("open rabbitMQ channel failed, %v", err)
	}
	defer ch.Close()

	//declare consumer queue
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	if err != nil {
		c.GinContext.JSON(-1, gin.H{
			"error": "declare queue failed," + err.Error(),
		})
		return
	}

	//consume queue
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		c.GinContext.JSON(-1, gin.H{
			"error": "consume queue failed," + err.Error(),
		})
		return
	}
	sliceBuf := make([]string, 0, 32)
	//loop consume queue to get message
	//for d := range msgs {
	//	fmt.Println(string(d.Body))
	//	sliceBuf = append(sliceBuf, string(d.Body))
	//}

	go func() {
		for d := range msgs {
			sliceBuf = append(sliceBuf, string(d.Body))
		}
	}()

	c.GinContext.JSON(200, gin.H{
		"message": sliceBuf,
	})
}
