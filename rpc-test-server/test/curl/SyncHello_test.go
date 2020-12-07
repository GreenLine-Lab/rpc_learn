package curl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"rpc/rpc-test-server/pb"
	"testing"
	"time"
)

func TestSyncHello(t *testing.T) {

	message := pb.ReqTest{
		Name: "Vladimir",
	}

	rBody, _ := message.Descriptor()
	rUrl := "http://localhost:9050/proto.TestServer/TestSyncHello"

	client := http.Client{
		Transport: &http.Transport{
			ResponseHeaderTimeout: time.Second,
		},
	}

	buffer := bytes.NewBuffer(rBody)
	req, err := http.NewRequest("POST", rUrl, buffer)
	if err != nil {
		t.Fatalf("Unable create new HTTP request: %s", err.Error())
	}

	req.Header.Set("content-type", "application/protobuf")

	rpl, err := client.Do(req)
	if err != nil {
		t.Fatalf("client.Do: %s", err.Error())
	}

	body, err := ioutil.ReadAll(rpl.Body)
	if err != nil {
		t.Fatalf("Unable read response body: %s", err.Error())
	}

	defer func() {
		if err := rpl.Body.Close(); err != nil {
			t.Errorf("Unable close response body")
		}
	}()

	fmt.Printf("%s", string(body))
}
