package gopusher

import (
	"testing"
	"os"
	"fmt"
)


type config struct{
	appId string
	appKey string
	secretKey string
}

func Test_New_CreateNewPusher2(t *testing.T) {
	if os.Getenv("INTEGRATION") != "true"{
		return
	}

	c := config{os.Getenv("APPID"),os.Getenv("APPKEY"),os.Getenv("SECRETKEY")}

	pusher := New(c.appId, c.appKey, c.secretKey)

	result, statusCode, err := pusher.Trigger([]string{"test"}, "event", someData , nil)

	fmt.Println("result", result)
	fmt.Println("statusCode", statusCode)
	fmt.Println("err", err)

	if statusCode != 200 {
		t.Error("expected statuscode 200, actual ", statusCode)
	}

}
