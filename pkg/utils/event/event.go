package event

import (
	"encoding/json"
)

type Event struct {
	PartitionKey string         `json:"partitionKey"`
	Scheme       string         `json:"scheme"`
	EventId      string         `json:"eventId"`
	Timestamp    int            `json:"timestamp"`
	Stage        string         `json:"stage"`
	Data         map[string]any `json:"data"`
}

func ConvertByteArrayToEvent(eventEncoded []byte) (Event, error) {
	var event Event
	err := json.Unmarshal(eventEncoded, &event)

	return event, err
}
