package httpclient

import (
	"io"
	"net/http"

	"github.com/devMarcus21/eventfunnel/pkg/utils/responses"
)

type HttpClientWrapper struct{}

func (*HttpClientWrapper) Send(request *http.Request) (responses.HttpResponse, error) {
	// send the request
	client := &http.Client{}
	response, sendRequestError := client.Do(request)

	if sendRequestError != nil {
		// TODO log this error at some point once system wide logging is implemented
		return responses.HttpResponse{}, sendRequestError
	}

	responseBody, readResponseBodyError := io.ReadAll(response.Body)
	defer response.Body.Close()
	if readResponseBodyError != nil {
		// TODO log this at some point
		return responses.HttpResponse{}, readResponseBodyError
	}

	return responses.HttpResponse{ResultCode: response.StatusCode, BodyResult: responseBody}, nil
}
