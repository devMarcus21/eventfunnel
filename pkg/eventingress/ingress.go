package main

import (
	"log"
	"net/http"

	ingressfactory "github.com/devMarcus21/eventfunnel/pkg/eventingress/eventingressfactory"
	"github.com/devMarcus21/eventfunnel/pkg/utils/dbutils/model"
	"github.com/devMarcus21/eventfunnel/pkg/utils/httpwrapper"
)

func main() {
	modelTable := model.GetModelTable()

	http.HandleFunc("/ingress", httpwrapper.BuildHttpHandler(ingressfactory.CreateEventIngressHandler(modelTable)))

	log.Fatal(http.ListenAndServe(":80", nil))
}
