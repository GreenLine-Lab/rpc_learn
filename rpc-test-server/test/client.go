package test

import (
	"google.golang.org/grpc"
	"rpc-learn/rpc-test-server/pb"
	"sync"
)

const defaultUrl = ":9050"

type TestClient struct {
	mux        sync.Mutex
	connection *grpc.ClientConn
	url        string
}

func (c *TestClient) Client() pb.TestServerClient {

	if c.connection != nil {
		if err := c.connection.Close(); err != nil {
			return nil
		}
	} else {
		c.connection = &grpc.ClientConn{}
	}

	if len(c.url) == 0 {
		c.SetUrl()
	}

	var err error
	c.connection, err = grpc.Dial(c.url, grpc.WithInsecure())
	if err != nil {
		return nil
	}

	return pb.NewTestServerClient(c.connection)
}

func (c *TestClient) SetUrl(url ...string) {
	c.url = defaultUrl

	if url != nil {
		if len(url[0]) != 0 {
			c.url = url[0]
		}
	}
}

func (c *TestClient) CloseConnection() {
	if c.connection != nil {
		_ = c.connection.Close()
	}
}
