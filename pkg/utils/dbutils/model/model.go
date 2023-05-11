package model

type Model struct {
	Name   string         `json:"name"`
	Stage  string         `json:"stage"`
	Scheme map[string]any `json:"scheme"` // saved in cassandra models table as a string
}
