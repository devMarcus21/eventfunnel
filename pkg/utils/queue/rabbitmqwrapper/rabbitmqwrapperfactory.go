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
func GetRabbitmqPublisherWrapper(
	queueServiceConnection string,
	queueName string,
	createRabbitmqConnection func(queueServiceConnectionString string, queueName string) (RabbitmqConnections, error),
) func(event.Event) bool {
	// TODO log error if it exists/retry
	return func(event event.Event) bool {
		rabbitmqConnection, err := createRabbitmqConnection(queueServiceConnection, queueName)
		if err != nil {
			return false
		}
		defer rabbitmqConnection.amqpConnection.Close()
		defer rabbitmqConnection.amqpChannel.Close()

		q, qErr := rabbitmqConnection.amqpChannel.QueueDeclare(
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
		publishErr := rabbitmqConnection.amqpChannel.PublishWithContext(ctx,
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
