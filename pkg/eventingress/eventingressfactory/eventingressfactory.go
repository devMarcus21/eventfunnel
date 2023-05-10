package eventingress

import (
	"github.com/devMarcus21/eventfunnel/pkg/utils/dbutils/model"
	e "github.com/devMarcus21/eventfunnel/pkg/utils/event"
	res "github.com/devMarcus21/eventfunnel/pkg/utils/responses"
)

func CreateEventIngressHandler(modelTable func(string, string) model.Model) func(e.Event) res.ServiceResponse {
	return func(event e.Event) res.ServiceResponse {
		model := modelTable(event.Model, event.Stage)

		return res.ServiceResponse{
			"status":  "sucess",
			"model":   model,
			"payload": event,
		}
	}
}
