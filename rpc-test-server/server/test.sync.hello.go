package server

import (
	"context"
	"rpc/rpc-test-server/pb"
)

func (srv *TestServer) TestSyncHello(ctx context.Context, req *pb.ReqTest) (*pb.RplTest, error) {
	return &pb.RplTest{
		Message: "Hello, " + req.Name + "!",
	}, nil
}
