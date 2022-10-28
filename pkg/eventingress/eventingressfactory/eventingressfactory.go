package eventingress

import (
	e "github.com/devMarcus21/eventfunnel/pkg/utils/event"
	res "github.com/devMarcus21/eventfunnel/pkg/utils/responses"
)

func CreateEventIngressHandler() func(e.Event) res.ServiceResponse {
	return func(event e.Event) res.ServiceResponse {
		return res.ServiceResponse{
			"status":  "sucess",
			"payload": event,
		}
	}
}
