package test

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
	defer c.CloseConnection()

	c.mux.Lock()
	rpl, err := c.Client().TestSyncHello(context.Background(), &req)
	c.mux.Unlock()

	if err != nil {
		t.Fatalf("TestSyncHello return: %s", err.Error())
	}

	fmt.Printf("Rpl: %+v\n", rpl)
}
