package main

import (
	"log"
	"net/http"

	ingressfactory "github.com/devMarcus21/eventfunnel/pkg/eventingress/eventingressfactory"
	"github.com/devMarcus21/eventfunnel/pkg/utils/dbutils/model"
	"github.com/devMarcus21/eventfunnel/pkg/utils/httpwrapper"
	"github.com/devMarcus21/eventfunnel/pkg/utils/schemevalidation"
)

func main() {
	modelTable := model.GetModelTable()
	schemevalidatior := schemevalidation.GetSchemeValidator()

	http.HandleFunc("/ingress", httpwrapper.BuildHttpHandler(ingressfactory.CreateEventIngressHandler(modelTable, schemevalidatior)))

	log.Fatal(http.ListenAndServe(":80", nil))
}
