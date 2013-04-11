package server

import (
	"net/http"
  "io/ioutil"
  "bytes"
)

type Pusher struct{
	appId string
	appKey string
	appSecret string
	poster Poster
}

func CreatePusher() Pusher{
	return Pusher{poster : PusherHttpEndpoint{}}
}
//error if keys are not set
//  when? On instantiate or on Trigger()

//post
//get
//auth
//createSignedQueryString

type Poster interface{
	post(body []byte)(data *string, err error)
}

type PusherHttpEndpoint struct{
	url string
}

func (this PusherHttpEndpoint) post(data []byte)(*string, error){
	httpclient := &http.Client{}
	req, err := http.NewRequest("POST", this.url, bytes.NewBuffer(data))
	resp, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	body := string(responseData)
	return &body, err
}

func (p Pusher) Trigger(channel, event string, data interface{}) string {
	//data collect and parse
	//post


	return "ok"
}

