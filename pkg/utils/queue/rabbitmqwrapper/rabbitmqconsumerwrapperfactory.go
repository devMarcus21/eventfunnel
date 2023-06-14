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
			numberOfMessagesInChannel := len(msgs) // number of messages currently in the rabbitmq channel

			for i := 0; i < numberOfMessagesInChannel && i < 10; i++ {
				/*select {
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
						fmt.Println("NOT OK")
						break // channel is closed
					}
				default:
					fmt.Println("default break")
					break // no value in channel
				}*/
				msg := <-msgs
				event, err := event.ConvertByteArrayToEvent(msg.Body)
				if err != nil {
					// TODO log marshalling error
					fmt.Println("FAILED CONVERSION")
					continue
				}
				events = append(events, event)
				messageAckError := msg.Ack(false)
				if messageAckError != nil {
					// TODO dead code, just reminder to later log this error
					fmt.Println("FALIED ACK")
					continue
				}
			}

			if len(events) > 0 {
				fmt.Println(events)
			}
		}
	}
}
