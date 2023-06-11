package main

import (
	"fmt"
	"os"

	queueFactory "github.com/devMarcus21/eventfunnel/pkg/utils/queue"
)

func main() {
	podName := os.Getenv("POD_NAME") // ingress-consumer-{ordinal index}
	ordinal := podName[23:]          // get ordinal
	queueName := "ingress-queue-" + ordinal
	//fmt.Println("HELLO")
	fmt.Println(podName)
	fmt.Println(ordinal)
	fmt.Println(queueName)

	fmt.Println("amqp://guest:guest@" + queueName + ":5672/")
	consumerHandler := queueFactory.GetQueueConsumerWrapper("amqp://guest:guest@"+queueName+":5672/", "ingress-queue")

	// TODO log error if generated and kill this pod
	err := consumerHandler() // run the application code
	fmt.Println(err)
}
