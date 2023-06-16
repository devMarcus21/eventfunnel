package httpclient

import (
	"net/http"

	"github.com/devMarcus21/eventfunnel/pkg/utils/responses"
)

type HttpClientSender interface {
	Send(*http.Request) (responses.HttpResponse, error)
}
