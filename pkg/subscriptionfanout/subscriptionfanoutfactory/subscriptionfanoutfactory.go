package subscriptionfanoutfactory

import (
	"github.com/devMarcus21/eventfunnel/pkg/utils/event"
	"github.com/devMarcus21/eventfunnel/pkg/utils/responses"
)

func CreateSubscriptionFanoutHandler() func([]event.Event) responses.ServiceResponse {
	return func(events []event.Event) responses.ServiceResponse {
		return responses.ServiceResponse{
			"Status":  "Success",
			"Payload": events,
		}
	}
}
