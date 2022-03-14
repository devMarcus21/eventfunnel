package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventCreation(t *testing.T) {
	eventString := `{"itemKey":"item-key","schema":"schema","eventKey":"event-key","data":{"key1":"123","key2":123,"key3":false}}`

	event := CreateEvent(eventString)
	eventData := make(map[string]interface{})
	eventData["key1"] = "123"
	eventData["key2"] = float64(123)
	eventData["key3"] = false

	type expected struct {
		itemKey  string
		schema   string
		eventKey string
		data     map[string]interface{}
	}

	expectedEvent := expected{"item-key", "schema", "event-key", eventData}

	assert.Equal(t, expectedEvent.itemKey, event.GetItemKey(), "itemKey should be equal")
	assert.Equal(t, expectedEvent.schema, event.GetSchema(), "schema should be equal")
	assert.Equal(t, expectedEvent.eventKey, event.GetEventKey(), "eventKey should be equal")
	assert.Equal(t, expectedEvent.data, event.GetData(), "event data should be equal")
}
