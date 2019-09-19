package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//func Something() {
//	Post("https://api.github.com/user/repos", `{"name":"golang-turorial"}`)
//}

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBytes))
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}
