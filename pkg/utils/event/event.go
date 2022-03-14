package event

import (
	"encoding/json"
)

type Event struct {
	ItemKey  string
	Schema   string
	EventKey string
	Data     map[string]interface{}
}

func (e Event) GetItemKey() string {
	return e.ItemKey
}

func (e Event) GetSchema() string {
	return e.Schema
}

func (e Event) GetEventKey() string {
	return e.EventKey
}

func (e Event) GetData() map[string]interface{} {
	return e.Data
}

func CreateEvent(eventJson string) *Event {
	var event Event
	json.Unmarshal([]byte(eventJson), &event)
	return &event
}
