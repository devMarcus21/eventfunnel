package rabbitmqwrapper

import (
	"fmt"

	"github.com/devMarcus21/eventfunnel/pkg/utils/event"
)

func GetRabbitmqConsumerWrapper(
	queueServiceConnection string,
	queueName string,
	createRabbitmqConnection func(queueServiceConnectionString string, queueName string) (RabbitmqConnections, error),
) func() error {
	return func() error {
		rabbitmqConnection, err := createRabbitmqConnection(queueServiceConnection, queueName)
		if err != nil {
			return err
		}
		defer rabbitmqConnection.amqpConnection.Close()
		defer rabbitmqConnection.amqpChannel.Close()

		// consume
		msgs, consumeError := rabbitmqConnection.amqpChannel.Consume(
			rabbitmqConnection.queue.Name,
			"",
			false,
			false,
			false,
			false,
			nil,
		)
		if consumeError != nil {
			return consumeError
		}

		for {
			events := []event.Event{}

			// TODO make this a passed in env variable
			for i := 0; i < 10; i++ { // TODO maybe use range to make this cleaner
				select { // TODO make this logic cleaner. Make this more batchable
				case msg, ok := <-msgs:
					if ok {
						event, err := event.ConvertByteArrayToEvent(msg.Body)
						if err != nil {
							// TODO log marshalling error
							continue
						}
						events = append(events, event)
						messageAckError := msg.Ack(false)
						if messageAckError != nil {
							// TODO dead code, just reminder to later log this error
							continue
						}
					} else {
						break // channel is closed
					}
				default:
					break // no value in channel
				}
			}

			if len(events) > 0 {
				fmt.Println(events)
			}
		}
	}
}
