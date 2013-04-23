package gopusher

import (
	"testing"
	"strings"
)

var (
	aChannel = "mychannel"
	aEvent   = "gopher-stolen"
	someData = struct {
		Name    string
		Company string
	}{
		"Pieter Joost van de Sande",
		"Wercker",
	}
)

type mockEndpoint struct{
	postPath string
	postBody []byte

}

func (this *mockEndpoint) post(path string, body []byte) (data string, statusCode int, err error) {
	this.postPath = path
	this.postBody = body
	resultBody := "{\"ok\" : true}"
	return resultBody, 0, nil
}


func Test_New_CreateNewPusher(t *testing.T) {
	pusher := New("id", "key", "secret")
	if pusher == nil {
		t.Error("pusher should be created")
	}
}

// id, key and secret should be set
// pusher options:  port and https (restclient??)

func Test_Trigger_ShouldPostRequestToServer(t *testing.T){
	mockEndpoint := &mockEndpoint{}
	pusher := pusher{"id", "key", "secret", mockEndpoint}

	body, _, _ := pusher.Trigger([]string{"channel"}, "event", someData, nil)

	if(body == ""){
		t.Error("Body should be filled")
	}
}

func Test_Trigger_ShouldPostCorrectlyFormattedRequestToServer(t *testing.T){
	mockEndpoint := &mockEndpoint{}
	pusher := pusher{"id", "key", "secret", mockEndpoint}

	pusher.Trigger([]string{"channel"}, "event", someData, nil)

	if !strings.HasPrefix(mockEndpoint.postPath, "/apps/id/events"){
		t.Error("path not set correctly ", mockEndpoint.postPath)
	}
	expectedBody := "{\"channels\":[\"channel\"],\"name\":\"event\",\"data\":\"{\\\"Name\\\":\\\"Pieter Joost van de Sande\\\",\\\"Company\\\":\\\"Wercker\\\"}\"}"
	if string(mockEndpoint.postBody) != expectedBody {
		t.Errorf("body not set correctly, expected %v, actual %v", expectedBody, string(mockEndpoint.postBody))
	}
}


//error on no name and/or data

// check default https
//filter null json
// one channel: use channel, else use channels


