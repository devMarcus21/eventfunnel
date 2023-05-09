package model

type model struct {
	Name        string         `json:"name"`
	Stage       string         `json:"stage"`
	ModelScheme map[string]any `json:"modelScheme"`
}
