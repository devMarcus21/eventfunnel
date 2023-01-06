package events

type Event struct {
	PartitionKey string         `json:"partitionKey"`
	Schema       string         `json:"schema"`
	EventId      string         `json:"eventId"`
	Time         int64          `json:"time"`
	Data         map[string]any `json:"data"`
}
