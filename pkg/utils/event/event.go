package event

type Event struct {
	PartitionKey string
	Schema       string
	EventId      string
	Time         int64
}
