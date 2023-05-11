package model

import (
	"encoding/json"
)

func deserializeMapToModel(model map[string]any) Model {
	// saved in cassandra models table as a string so need to convert it to proper type
	var scheme map[string]any
	// TODO handle errors
	json.Unmarshal([]byte(model["scheme"].(string)), &scheme)

	// TODO handle type checking here
	return Model{
		Name:   model["name"].(string),
		Stage:  model["stage"].(string),
		Scheme: scheme,
	}
}

func CreateModelTable(cassandra func(params ...interface{}) []map[string]any) func(string, string) Model {
	return func(modeName string, stage string) Model {
		queryResults := cassandra(modeName, stage)

		// if query results is empty, no schema was found for the given stage and schema
		if len(queryResults) == 0 {
			return Model{}
		}

		// Only interested in first instance in array
		// There should not be more than one model for model name + stage
		return deserializeMapToModel(queryResults[0])
	}
}
