package gopusher

import (
	"net/http"
	"strings"
	"io/ioutil"
)


 // domain: 'api.pusherapp.com',
 //    scheme: 'http',
 //    port: 80

type endpointer interface{
	post(path string, body []byte) (data string, statusCode int, err error)
}

type endpoint struct{

}

func (e *endpoint) post(path string, body []byte) (data string, statusCode int, err error) {
	url := "http://api.pusherapp.com" + path
	bodyReader := strings.NewReader(string(body))
	res, err := http.Post(url, "application/json", bodyReader)

	if err != nil {
		return "", 0, err
	}
	responseBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", 0, err
	}

	return string(responseBody), res.StatusCode, nil
}

func newEndpoint() endpointer{
	return &endpoint{}
}

