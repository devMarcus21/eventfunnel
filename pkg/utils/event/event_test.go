package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventCreation(t *testing.T) {
	eventString := `{"itemKey":"item-key","schema":"schema","eventKey":"event-key","data":{"key1":"123","key2":123,"key3":false}}`

	event := CreateEvent(eventString)
	itemKey := "item-key"
	schema := "schema"
	eventKey := "event-key"

	assert.Equal(t, itemKey, event.GetItemKey(), "itemKey should be equal")
	assert.Equal(t, schema, event.GetSchema(), "schema should be equal")
	assert.Equal(t, eventKey, event.GetEventKey(), "eventKey should be equal")
}
