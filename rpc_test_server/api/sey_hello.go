package api

import (
	"context"
	"fmt"
	"rpc_learn/lib/zlog"
	pb2 "rpc_learn/rpc_test_server/pkg/pb"
)

func (srv *TestServer) TestSeyHello(ctx context.Context, req *pb2.ReqTestSeyHello) (*pb2.RplTestSeyHello, error) {
	rpl := &pb2.RplTestSeyHello{}
	log := zlog.ApiLogger(srv.logger, ctx)

	log.Trace().Msg("Start processing")
	defer log.Trace().Msg("Finish processing")

	rpl.Message = fmt.Sprintf("Hello, %s!", req.Name)

	return rpl, nil
}
