package main

import (
	"log"
	"net/http"

	"github.com/devMarcus21/eventfunnel/pkg/subscriptionfanout/subscriptionfanoutfactory"
	"github.com/devMarcus21/eventfunnel/pkg/utils/httpwrapper"
)

func main() {
	http.HandleFunc("/subscriptionfanout", httpwrapper.BuildHttpHandlerBatch(subscriptionfanoutfactory.CreateSubscriptionFanoutHandler()))
	log.Fatal(http.ListenAndServe(":80", nil))
}
