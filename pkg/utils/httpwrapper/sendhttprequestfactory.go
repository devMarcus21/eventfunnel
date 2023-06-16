package httpwrapper

import (
	"bytes"
	"net/http"

	"github.com/devMarcus21/eventfunnel/pkg/utils/httpwrapper/httpclient"
	"github.com/devMarcus21/eventfunnel/pkg/utils/responses"
)

func CreateHttpClientRequestWrapper(httpClient httpclient.HttpClientSender, url string, method string) func([]byte) (responses.HttpResponse, error) {
	return func(message []byte) (responses.HttpResponse, error) {
		request, buildingRequestError := http.NewRequest(method, url, bytes.NewBuffer(message))
		if buildingRequestError != nil {
			return responses.HttpResponse{}, buildingRequestError
		}
		request.Header.Set("Content-Type", "application/json; charset=utf-8")

		return httpClient.Send(request)
	}
}
