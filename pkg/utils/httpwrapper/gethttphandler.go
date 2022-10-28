package httpwrapper

import (
	"net/http"
)

func BuildHttpHandler(f func() string) func(w http.ResponseWriter, r *http.Request) {
	response := f()
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
	}
}
