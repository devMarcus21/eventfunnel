package eventingress

import (
	"github.com/devMarcus21/eventfunnel/pkg/utils/dbutils/scheme"
	"github.com/devMarcus21/eventfunnel/pkg/utils/event"
	res "github.com/devMarcus21/eventfunnel/pkg/utils/responses"
)

func CreateEventIngressHandler(
	getSchemeFromTable func(string, string) (scheme.Scheme, error),
	schemeValidator func(event.Event, scheme.Scheme) bool,
	queuePublisher func(event.Event) bool,
) func(event.Event) res.ServiceResponse {

	return func(event event.Event) res.ServiceResponse {
		scheme, err := getSchemeFromTable(event.Scheme, event.Stage)
		if err != nil {
			return res.ServiceResponse{
				"status":  "failed",
				"payload": event,
				"message": err.Error(),
			}
		}
		validScheme := schemeValidator(event, scheme)

		if !validScheme {
			return res.ServiceResponse{
				"status":  "failed",
				"payload": event,
				"message": "Scheme is invalid or doesn't exist",
			}
		}

		if !queuePublisher(event) {
			return res.ServiceResponse{
				"status":  "failed",
				"payload": event,
				"message": "event failed to publish to ingress queue",
			}
		}

		return res.ServiceResponse{
			"status":           "sucess",
			"scheme":           scheme,
			"payload":          event,
			"isValidScheme":    true,
			"payloadPublished": true,
		}
	}
}
