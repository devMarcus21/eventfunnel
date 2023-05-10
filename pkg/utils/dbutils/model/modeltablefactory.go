package model

import (
	"encoding/json"
)

func deserializeMapToModel(model map[string]any) Model {
	var serializedModel Model

	// TODO handle possible errors
	jsonData, _ := json.Marshal(model)
	json.Unmarshal(jsonData, &serializedModel)

	return serializedModel
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
