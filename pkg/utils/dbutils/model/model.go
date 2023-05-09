package model

type model struct {
	Name   string         `json:"name"`
	Stage  string         `json:"stage"`
	Scheme map[string]any `json:"scheme"`
}
