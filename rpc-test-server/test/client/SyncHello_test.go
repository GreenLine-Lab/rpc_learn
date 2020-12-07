package client

import (
	"context"
	"fmt"
	"rpc/rpc-test-server/pb"
	"testing"
)

func TestHelloSync(t *testing.T) {

	req := pb.ReqTest{
		Name: "Vladimir",
	}

	c := TestClient{}
	if err := c.InitClient("localhost:9050"); err != nil {
		t.Fatalf("NewConnection: %s", err.Error())
	}
	defer c.Connection.Close()

	rpl, err := c.Client.TestSyncHello(context.Background(), &req)
	if err != nil {
		t.Fatalf("TestSyncHello return: %s", err.Error())
	}

	fmt.Println(rpl.Message)
}
