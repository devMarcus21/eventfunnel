package scheme

type Scheme struct {
	Name  string         `json:"name"`
	Stage string         `json:"stage"`
	Data  map[string]any `json:"data"` // saved in cassandra scheme table as a string
}
