package eventingress

import (
	"github.com/devMarcus21/eventfunnel/pkg/utils/dbutils/model"
	e "github.com/devMarcus21/eventfunnel/pkg/utils/event"
	res "github.com/devMarcus21/eventfunnel/pkg/utils/responses"
)

func CreateEventIngressHandler(getModelFromTable func(string, string) model.Model, schemeValidator func(e.Event, model.Model) bool) func(e.Event) res.ServiceResponse {
	return func(event e.Event) res.ServiceResponse {
		model := getModelFromTable(event.Model, event.Stage)
		validateScheme := schemeValidator(event, model)

		return res.ServiceResponse{
			"status":        "sucess",
			"model":         model,
			"payload":       event,
			"isValidScheme": validateScheme,
		}
	}
}
