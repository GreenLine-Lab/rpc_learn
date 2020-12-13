package api

import (
	"context"
	"fmt"
	"rpc-learn/lib/zlog"
	"rpc-learn/rpc-test-server/pb"
)

func (srv *TestServer) TestSeyHello(ctx context.Context, req *pb.ReqTestSeyHello) (*pb.RplTestSeyHello, error) {
	rpl := &pb.RplTestSeyHello{}
	log := zlog.ApiLogger(srv.logger, ctx)

	log.Trace().Msg("Start processing")
	defer log.Trace().Msg("Finish processing")

	rpl.Message = fmt.Sprintf("Hello, %s!", req.Name)

	return rpl, nil
}
