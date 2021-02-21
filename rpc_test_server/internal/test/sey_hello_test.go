package test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"rpc-learn/rpc_test_server/pb"
	"testing"
	"time"
)

func TestHello(t *testing.T) {

	req := pb.ReqTestSeyHello{
		Name: "Vladimir",
	}

	c := NewRpcTestServerClient()
	defer c.ConnClose()

	reqMD := metadata.MD{}
	reqMD.Set("Trace-Id", c.GetTraceId())
	reqCtx := metadata.NewOutgoingContext(context.Background(), reqMD)

	rplMD := metadata.MD{}
	rpl, err := c.Client().TestSeyHello(reqCtx, &req, grpc.Header(&rplMD))
	if err != nil {
		t.Fatalf("TestSyncHello return: %s", err.Error())
	}

	fmt.Printf("%s [%s] Rpl: %+v\n", time.Now().Format(time.StampNano), c.GetTraceId(), rpl)
}
