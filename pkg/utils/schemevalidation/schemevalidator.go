package schemevalidation

import (
	"reflect"

	"github.com/devMarcus21/eventfunnel/pkg/utils/dbutils/scheme"
	"github.com/devMarcus21/eventfunnel/pkg/utils/event"
)

func deepEqualsCompare(eventData map[string]any, schemeData map[string]any) bool {
	for schemeKey, schemeVal := range schemeData {
		eventVal, exists := eventData[schemeKey]
		if !exists {
			return false
		}

		eventValTypeString := reflect.TypeOf(eventVal).String()
		schemaValTypeString := reflect.TypeOf(schemeVal).String()
		if eventValTypeString == "map[string]interface {}" || schemaValTypeString == "map[string]interface {}" {
			// One or both of these are of type map[string]any so check for type comparison (eventValTypeString == schemaValTypeString)
			if eventValTypeString != "map[string]interface {}" || schemaValTypeString != "map[string]interface {}" {
				return false
			}

			if !deepEqualsCompare(eventVal.(map[string]any), schemeVal.(map[string]any)) {
				return false
			}
			continue
		}

		if eventValTypeString != schemeVal.(string) {
			return false
		}
	}
	return true
}

func GetSchemeValidator() func(event.Event, scheme.Scheme) bool {
	return func(event event.Event, scheme scheme.Scheme) bool {
		// TODO validation will return true for event that has extra fields. Determine if this is appropriate
		return deepEqualsCompare(event.Data, scheme.Data)
	}
}
