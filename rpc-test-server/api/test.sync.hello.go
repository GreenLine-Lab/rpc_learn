package api

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"rpc-learn/rpc-test-server/pb"
)

func (srv *TestServer) TestSyncHello(ctx context.Context, req *pb.ReqTest) (*pb.RplTest, error) {
	rpl := &pb.RplTest{}
	method, _ := grpc.Method(ctx)
	log.Println(method)

	rpl.Message = fmt.Sprintf("Hello, %s!", req.Name)

	return rpl, nil
}
