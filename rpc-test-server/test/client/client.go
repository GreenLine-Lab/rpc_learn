package client

import (
	"google.golang.org/grpc"
	"rpc/rpc-test-server/pb"
)

type TestClient struct {
	Connection *grpc.ClientConn
	Client     pb.TestServerClient
}

func (c *TestClient) NewConnection(url string) error {
	var err error
	c.Connection, err = grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return err
	}

	return nil
}

func (c *TestClient) NewClient() {
	c.Client = pb.NewTestServerClient(c.Connection)
}

func NewTestClient() TestClient {
	return TestClient{}
}
