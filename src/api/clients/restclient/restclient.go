package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

//func Something() {
//	Post("https://api.github.com/user/repos", `{"name":"golang-turorial"}`)
//}

var (
	enabledMocks = false
	mocks        = make(map[string]*Mock)
)

type Mock struct {
	Url        string
	HttpMethod string
	Response   *http.Response
	Err        error
}

func getMockId(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

func StartMockups() {
	enabledMocks = true
}

func FlushMockups() {
	mocks = make(map[string]*Mock)
}

func StopMockups() {
	enabledMocks = false

}

func AddMockup(mock Mock) {
	mocks[getMockId(mock.HttpMethod, mock.Url)] = &mock
}

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enabledMocks {
		mock := mocks[getMockId(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("no mockup found for give request")
		}
		return mock.Response, mock.Err
		// return llcaol mock without calling any external resource
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBytes))
	request.Header = headers
	log.Println("##req", request.Header, request.Body)

	client := http.Client{}
	return client.Do(request)
}
