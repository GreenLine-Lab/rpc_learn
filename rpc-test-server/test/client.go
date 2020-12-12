package test

import (
	"google.golang.org/grpc"
	"log"
	"rpc-learn/rpc-test-server/pb"
)

const defaultUrl = ":9050"

type Client struct {
	connection *grpc.ClientConn
	url        string
}

func (c *Client) Client() pb.TestServerClient {

	if c.connection != nil {
		if err := c.connection.Close(); err != nil {
			log.Printf("Conn close: %s", err.Error())
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

func (c *Client) SetUrl(url ...string) {
	c.url = defaultUrl

	if url != nil {
		if len(url[0]) != 0 {
			c.url = url[0]
		}
	}
}

func (c *Client) CloseConnection() {
	if c.connection != nil {
		_ = c.connection.Close()
	}
}
