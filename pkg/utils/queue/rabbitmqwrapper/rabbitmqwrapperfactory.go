package rabbitmqwrapper

import (
	"context"
	"encoding/json"
	"time"

	"github.com/devMarcus21/eventfunnel/pkg/utils/event"

	amqp "github.com/rabbitmq/amqp091-go"
)

// TODO make this more modular
// TODO look at connection pooling
func GetRabbitmqPublisherWrapper(queueServiceConnection string, queueName string) func(event.Event) bool {
	// TODO log error if it exists/retry
	return func(event event.Event) bool {
		conn, connectionErr := amqp.Dial(queueServiceConnection)
		if connectionErr != nil {
			return false
		}
		defer conn.Close()

		ch, channelErr := conn.Channel()
		if channelErr != nil {
			return false
		}
		defer ch.Close()
		// TODO look at adding Qos control github.com/rabbitmq/amqp091-go/blob/main/channel.go

		q, qErr := ch.QueueDeclare(
			queueName, // name
			false,     // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)
		if qErr != nil {
			return false
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		eventEncoded, encodeErr := json.Marshal(event)
		if encodeErr != nil {
			return false
		}
		publishErr := ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        eventEncoded,
			})
		if publishErr != nil {
			return false
		}

		return true
	}
}
