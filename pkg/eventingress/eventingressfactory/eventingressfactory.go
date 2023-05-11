package eventingress

import (
	"github.com/devMarcus21/eventfunnel/pkg/utils/dbutils/scheme"
	e "github.com/devMarcus21/eventfunnel/pkg/utils/event"
	res "github.com/devMarcus21/eventfunnel/pkg/utils/responses"
)

func CreateEventIngressHandler(getSchemeromTable func(string, string) scheme.Scheme, schemeValidator func(e.Event, scheme.Scheme) bool) func(e.Event) res.ServiceResponse {
	return func(event e.Event) res.ServiceResponse {
		scheme := getSchemeromTable(event.Scheme, event.Stage)
		validateScheme := schemeValidator(event, scheme)

		return res.ServiceResponse{
			"status":        "sucess",
			"scheme":        scheme,
			"payload":       event,
			"isValidScheme": validateScheme,
		}
	}
}
