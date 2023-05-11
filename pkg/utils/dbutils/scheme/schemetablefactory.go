package scheme

import (
	"encoding/json"
	"errors"
)

func deserializeMapToScheme(scheme map[string]any) Scheme {
	// saved in cassandra scheme table as a string so need to convert it to proper type
	var schema map[string]any
	// TODO handle errors
	json.Unmarshal([]byte(scheme["data"].(string)), &schema)

	// TODO handle type checking here
	return Scheme{
		Name:  scheme["name"].(string),
		Stage: scheme["stage"].(string),
		Data:  schema,
	}
}

func CreateSchemeTable(cassandra func(params ...interface{}) []map[string]any) func(string, string) (Scheme, error) {
	return func(modeName string, stage string) (Scheme, error) {
		queryResults := cassandra(modeName, stage)

		// if query results is empty, no schema was found for the given stage and schema
		if len(queryResults) == 0 {
			return Scheme{}, errors.New("Scheme not found in scheme table")
		}

		// Only interested in first instance in array
		// There should not be more than one scheme for scheme name + stage
		return deserializeMapToScheme(queryResults[0]), nil
	}
}
