package queue

import (
	"github.com/devMarcus21/eventfunnel/pkg/utils/event"
	rabbitmqWrapperFactory "github.com/devMarcus21/eventfunnel/pkg/utils/queue/rabbitmqwrapper"
)

// TODO add function batching support for this client. Rabbitmq itself may not support batching
func GetQueuePublisherWrapper(queueServiceConnection string, queueName string) func(event.Event) (bool, error) {
	return rabbitmqWrapperFactory.GetRabbitmqPublisherWrapper(queueServiceConnection, queueName, rabbitmqWrapperFactory.GetRabbitmqConnections)
}

func GetQueueConsumerWrapper(queueServiceConnection string, queueName string) func() error {
	return rabbitmqWrapperFactory.GetRabbitmqConsumerWrapper(queueServiceConnection, queueName, rabbitmqWrapperFactory.GetRabbitmqConnections)
}
