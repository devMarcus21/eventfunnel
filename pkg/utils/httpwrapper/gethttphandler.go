package httpwrapper

import (
	"encoding/json"
	"net/http"

	e "github.com/devMarcus21/eventfunnel/pkg/utils/event"
	res "github.com/devMarcus21/eventfunnel/pkg/utils/responses"
)

func BuildHttpHandler(handler func(e.Event) res.ServiceResponse) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var event e.Event

		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&event)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := handler(event)

		json.NewEncoder(w).Encode(response)
	}
}
