package main

import (
	"log"
	"net/http"

	ingressfactory "github.com/devMarcus21/eventfunnel/pkg/eventingress/eventingressfactory"
	"github.com/devMarcus21/eventfunnel/pkg/utils/dbutils/scheme"
	"github.com/devMarcus21/eventfunnel/pkg/utils/httpwrapper"
	"github.com/devMarcus21/eventfunnel/pkg/utils/queue"
	"github.com/devMarcus21/eventfunnel/pkg/utils/schemevalidation"
)

func main() {
	schemeTable := scheme.GetSchemeTable()
	schemevalidatior := schemevalidation.GetSchemeValidator()

	queuePublisher := queue.GetQueuePublisherWrapper("amqp://guest:guest@ingress-queue-service:5672/", "ingress-queue")

	http.HandleFunc("/ingress", httpwrapper.BuildHttpHandler(ingressfactory.CreateEventIngressHandler(schemeTable, schemevalidatior, queuePublisher)))

	log.Fatal(http.ListenAndServe(":80", nil))
}
