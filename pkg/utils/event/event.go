package event

import (
	"encoding/json"
	"fmt"
)

type Event map[string]interface{}

func (e Event) GetItemKey() string {
	return fmt.Sprintf("%v", e["itemKey"])
}

func (e Event) GetSchema() string {
	return fmt.Sprintf("%v", e["schema"])
}

func (e Event) GetEventKey() string {
	return fmt.Sprintf("%v", e["eventKey"])
}

func (e Event) GetData() interface{} {
	return e["data"]
}

func CreateEvent(eventJson string) *Event {
	var event Event
	json.Unmarshal([]byte(eventJson), &event)
	return &event
}
