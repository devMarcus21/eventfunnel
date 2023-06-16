package rabbitmqwrapper

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitmqConnections struct {
	amqpConnection *amqp.Connection
	amqpChannel    *amqp.Channel
	queue          amqp.Queue
}

func GetRabbitmqConnections(queueServiceConnectionString string, queueName string) (RabbitmqConnections, error) {
	conn, connectionErr := amqp.Dial(queueServiceConnectionString)
	if connectionErr != nil {
		return RabbitmqConnections{}, connectionErr
	}

	ch, channelErr := conn.Channel()
	if channelErr != nil {
		conn.Close()
		return RabbitmqConnections{}, channelErr
	}

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
		conn.Close()
		ch.Close()
		return RabbitmqConnections{}, qErr
	}

	return RabbitmqConnections{amqpConnection: conn, amqpChannel: ch, queue: q}, nil
}
