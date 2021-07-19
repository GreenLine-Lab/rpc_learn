package api

import (
	"context"
	pb2 "rpc_learn/rpc_test_server/pkg/pb"
)

func (srv *TestServer) TestDbUserCreate(ctx context.Context, req *pb2.ReqTestDbUserCreate) (*pb2.RplTestServer, error) {
	return nil, nil
}
