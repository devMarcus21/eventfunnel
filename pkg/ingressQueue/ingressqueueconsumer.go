package main

import (
	"fmt"

	queueFactory "github.com/devMarcus21/eventfunnel/pkg/utils/queue"
)

func main() {
	consumerHandler := queueFactory.GetQueueConsumerWrapper("amqp://guest:guest@localhost:5672/", "ingress-queue")

	// TODO log error if generated and kill this pod
	err := consumerHandler() // run the application code
	fmt.Println(err)         // TODO make this a log message
}
