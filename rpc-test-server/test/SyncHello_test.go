package test

import (
	"context"
	"fmt"
	"rpc-learn/rpc-test-server/pb"
	"testing"
)

func TestHelloSync(t *testing.T) {

	req := pb.ReqTest{
		Name: "Vladimir",
	}

	c := TestClient{}
	defer c.CloseConnection()

	rpl, err := c.Client().TestSyncHello(context.Background(), &req)
	if err != nil {
		t.Fatalf("TestSyncHello return: %s", err.Error())
	}

	fmt.Printf("Rpl: %+v\n", rpl)
}
