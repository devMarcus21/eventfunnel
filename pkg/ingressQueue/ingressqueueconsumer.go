package main

import (
	"fmt"

	"github.com/devMarcus21/eventfunnel/pkg/utils/httpwrapper"
	"github.com/devMarcus21/eventfunnel/pkg/utils/httpwrapper/httpclient"
	queueFactory "github.com/devMarcus21/eventfunnel/pkg/utils/queue"
)

func main() {
	subscriptionFanoutServiceUrl := "http://subscription-fanout-service.default.svc.cluster.local:3001/subscriptionfanout"

	httpClient := httpclient.HttpClientWrapper{}
	httpHandler := httpwrapper.CreateHttpClientRequestWrapper(httpClient, subscriptionFanoutServiceUrl, "POST")
	consumerHandler := queueFactory.GetQueueConsumerWrapper("amqp://guest:guest@localhost:5672/", "ingress-queue", httpHandler)

	// TODO log error if generated and kill this pod
	err := consumerHandler() // run the application code
	fmt.Println(err)         // TODO make this a log message
}
