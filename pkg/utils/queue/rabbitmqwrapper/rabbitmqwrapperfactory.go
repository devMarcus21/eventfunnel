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
) func(event.Event) (bool, error) {
	// TODO log error if it exists/retry
	return func(event event.Event) (bool, error) {
		rabbitmqConnection, err := createRabbitmqConnection(queueServiceConnection, queueName)
		if err != nil {
			return false, err
		}
		defer rabbitmqConnection.amqpConnection.Close()
		defer rabbitmqConnection.amqpChannel.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		eventEncoded, encodeErr := json.Marshal(event)
		if encodeErr != nil {
			return false, encodeErr
		}
		publishErr := rabbitmqConnection.amqpChannel.PublishWithContext(ctx,
			"",                            // exchange
			rabbitmqConnection.queue.Name, // routing key
			false,                         // mandatory
			false,                         // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        eventEncoded,
			})
		if publishErr != nil {
			return false, publishErr
		}

		return true, nil
	}
}
