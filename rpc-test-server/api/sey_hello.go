package api

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"rpc-learn/rpc-test-server/pb"
)

func (srv *TestServer) TestSeyHello(ctx context.Context, req *pb.ReqTestSeyHello) (*pb.RplTestSeyHello, error) {
	rpl := &pb.RplTestSeyHello{}

	method, _ := grpc.Method(ctx)
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		var traceId string
		if traceIdHeader := md.Get("Trace-Id"); traceIdHeader != nil {
			traceId = traceIdHeader[0]
		}

		log.Printf("%s [%s]: New request", method, traceId)
	}

	rpl.Message = fmt.Sprintf("Hello, %s!", req.Name)

	return rpl, nil
}
