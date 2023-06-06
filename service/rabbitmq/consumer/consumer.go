package consumer

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"log"
	"order-service/usecase"
)

var ListQueue = []string{
	"hello",
}

type Consumer struct {
	Conn    *amqp.Connection
	UseCase *usecase.UseCase
}

func NewConsumer(conn *amqp.Connection, uc *usecase.UseCase) *Consumer {
	return &Consumer{
		Conn:    conn,
		UseCase: uc,
	}
}

func (c *Consumer) StartConsumer() {
	for _, queue := range ListQueue {
		go c.Consume(queue)
	}
}

func (c *Consumer) Consume(queueName string) {
	ch, err := c.Conn.Channel()
	if err != nil {
		logrus.Error(err)
		return
	}
	defer ch.Close()
	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return
	}
	var forever chan struct{}

	go func() {
		for m := range msgs {
			switch queueName {
			case "hello":
				log.Println(m.Body)
			}
		}
	}()
	<-forever
	return
}
