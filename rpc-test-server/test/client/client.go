package client

import (
	"google.golang.org/grpc"
	"rpc/rpc-test-server/pb"
)

type TestClient struct {
	Connection *grpc.ClientConn
	Client     pb.TestServerClient
}

func (c *TestClient) InitClient(url string) error {
	var err error
	c.Connection, err = grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return err
	}

	c.Client = pb.NewTestServerClient(c.Connection)
	return nil
}
