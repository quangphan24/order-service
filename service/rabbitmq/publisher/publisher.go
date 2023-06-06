package publisher

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type Publisher struct {
	conn *amqp.Connection
}

func NewPublisher(conn *amqp.Connection) *Publisher {
	return &Publisher{conn: conn}
}
func (p *Publisher) Publish(queueName string, data interface{}) error {
	ch, err := p.conn.Channel()
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}
	body, _ := json.Marshal(data)
	if err := ch.PublishWithContext(context.Background(),
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		}); err != nil {
		return err
	}
	return nil
}
