package queue

import (
	"github.com/devMarcus21/eventfunnel/pkg/utils/event"
	rabbitmqwrapper "github.com/devMarcus21/eventfunnel/pkg/utils/queue/rabbitmqwrapper"
)

func GetQueuePublisherWrapper(queueServiceConnection string, queueName string) func(event.Event) bool {
	return rabbitmqwrapper.GetRabbitmqPublisherWrapper(queueServiceConnection, queueName)
}
