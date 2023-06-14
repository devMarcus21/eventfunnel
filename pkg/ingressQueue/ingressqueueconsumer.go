package main

import (
	"fmt"

	queueFactory "github.com/devMarcus21/eventfunnel/pkg/utils/queue"
)

func main() {
	fmt.Println("HELLO")

	//fmt.Println("amqp://guest:guest@" + queueName + ":5672/")
	//consumerHandler := queueFactory.GetQueueConsumerWrapper("amqp://guest:guest@"+queueName+":5672/", "ingress-queue")
	consumerHandler := queueFactory.GetQueueConsumerWrapper("amqp://guest:guest@localhost:5672/", "ingress-queue")

	// TODO log error if generated and kill this pod
	err := consumerHandler() // run the application code
	fmt.Println(err)
}
