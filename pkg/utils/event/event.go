package event

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	/*ItemKey  string
	Schema   string
	EventKey string
	Data     map[string]interface{}
	//Data string*/
	eventData map[string]interface{}
}

func (e *Event) GetItemKey() string {
	return fmt.Sprintf("%v", e.eventData["itemKey"])
}

func (e *Event) GetSchema() string {
	return fmt.Sprintf("%v", e.eventData["schema"])
}

func (e *Event) GetEventKey() string {
	return fmt.Sprintf("%v", e.eventData["eventKey"])
}

func (e *Event) GetData() interface{} {
	return e.eventData["data"]
}

func CreateEvent(eventJson string) *Event {
	var eventData map[string]interface{}
	json.Unmarshal([]byte(eventJson), &eventData)
	return &Event{eventData}
}
