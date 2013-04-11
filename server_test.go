package server

import ("testing")


type MockPusherEndpoint struct {
	postResult string
}

func (this MockPusherEndpoint) post(body []byte)(data *string, err error) {
	return &this.postResult, nil
}

func TestTesting(t *testing.T) {
	endpoint := MockPusherEndpoint{ postResult : "{iets:1}" }
	pusher := Pusher{poster : endpoint}
	result := pusher.Trigger("channel", "", "dus")

	if result != "ok" {
		t.Error("Das nie goe")
	}
}


