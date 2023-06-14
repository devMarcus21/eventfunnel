package rabbitmqwrapper

import (
	"fmt"

	"github.com/devMarcus21/eventfunnel/pkg/utils/event"
	amqp "github.com/rabbitmq/amqp091-go"
)

func GetRabbitmqConsumerWrapper(queueServiceConnection string, queueName string) func() error {
	return func() error {
		conn, connectionErr := amqp.Dial(queueServiceConnection)
		if connectionErr != nil {
			return connectionErr
		}
		defer conn.Close()

		ch, channelErr := conn.Channel()
		if channelErr != nil {
			return channelErr
		}
		defer ch.Close()

		q, qErr := ch.QueueDeclare(
			queueName, // name
			false,     // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)
		if qErr != nil {
			return qErr
		}

		// consume
		msgs, consumeError := ch.Consume(
			q.Name,
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
