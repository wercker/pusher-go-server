package gopusher

import (
	"testing"
	"time"
	//"fmt"
	)

func Test_createSignedQueryString_ShouldFailIfSystemParamsAreUsed(t *testing.T) {
	params := map[string]string { "auth_key" : "key"}
	body := "body"
	_, err := createSignedQueryString("key", time.Now(), &body, "path", "method", "secret", params)

	if err == nil{
		t.Error("err should not be nil")
	}
}

func Test_createSignedQueryString_ShouldReturnCorrectQueryString(t *testing.T) {
	key := "278d425bdf160c739803"
	secret := "7ad3773142a6692b25b8"
	body := "{\"name\":\"foo\",\"channels\":[\"project-3\"],\"data\":\"{\\\"some\\\":\\\"data\\\"}\"}"
	path :="/apps/3/events"
	method := "POST"
	timestamp := time.Unix(1353088179, 0)

	actual, err := createSignedQueryString(key, timestamp, &body, path, method, secret, nil)

	if err != nil{
		t.Error("err should be nil, ", err)
	}

	expected := "auth_key=278d425bdf160c739803&auth_timestamp=1353088179&auth_version=1.0&body_md5=ec365a775a4cd0599faeb73354201b6f&auth_signature=da454824c97ba181a32ccc17a72625ba02771f50b50e1e7430e47a1f3f457e6c"

	if actual != expected{
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}
