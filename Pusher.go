package gopusher

import (
	"fmt"
	"bytes"
	"encoding/json"
	"time"
	"crypto/md5"
	"errors"
	"sort"
	"crypto/hmac"
	"crypto/sha256"
)

type Pusher interface {
	Trigger(channels []string, event string, data interface{}, socketId *string)(body string, statusCode int, err error)
}

type pusher struct {
	appId string
	appKey string
	secretKey string
	endpoint endpointer
}


func (p *pusher) Trigger(channels []string, event string, data interface{}, socketId *string)(resultBody string, statusCode int, err error) {

	dataJson, err := json.Marshal(data)

	postData := struct {
		Channel *string `json:"channel,omitempty"`
		Channels *[]string `json:"channels,omitempty"`
		SocketId *string `json:"socket_id,omitempty"`
		Name string `json:"name"`
		Data  string `json:"data"`
	}{
		nil,
		&channels,
		nil,
		event,
		string(dataJson),
	}
	payloadData, err := json.Marshal(postData)

	if err != nil {
		return "", 0, err
	}
	payload := string(payloadData)

	queryString, err := createSignedQueryString(p.appKey, time.Now(), &payload, "/apps/" + p.appId + "/events", "POST", p.secretKey, nil)
	if err != nil{
		return "", 0, err
	}

	path := "/apps/" + p.appId + "/events?" + queryString

	return p.endpoint.post( path, payloadData)
	//return resultBody, nil, nil
}

func New(appId, appKey, secretKey string) Pusher {
	return &pusher{appId, appKey, secretKey, newEndpoint()}
}



func calcMd5Hash(data string) string {
	running_hash := md5.New(); // type hash.Hash
	running_hash.Write([]byte(data));  // data is []byte
	sum := running_hash.Sum(nil);
	return fmt.Sprintf("%x", sum)
}

func sortQueryString(params map[string]string) string{
	keys := make([]string,0)

	for key, _ := range params{
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var result bytes.Buffer
	first := true
	for _, key  := range keys{
		if first {
			first = false
		}	else {
			result.WriteString("&")
		}
		result.WriteString(key)
		result.WriteString("=")
		result.WriteString(params[key])
	}

	return result.String()
}

func createSignedQueryString(key string, timestamp time.Time, body *string, path, method, secret string, paramsQs map[string]string) (string, error){
	newParams := make(map [string]string)
	newParams["auth_key"] = key
	newParams["auth_timestamp"] = fmt.Sprint(timestamp.Unix())
	newParams["auth_version"] = "1.0"

	for key , _ := range paramsQs{
		if _,ok := newParams[key] ; ok {
			return "", errors.New(key + " is a required parameter and cannot be overidden")
		}
	}

	if body != nil{
		newParams["body_md5"] = calcMd5Hash(*body)
	}

	queryString := sortQueryString(newParams)

	signData := method + "\n" + path + "\n" + queryString

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(signData))
	sum := h.Sum(nil)

	result := queryString + "&auth_signature=" + fmt.Sprintf("%x", sum)

	return result, nil
}
