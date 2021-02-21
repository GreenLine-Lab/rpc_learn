package test

import (
	"github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"rpc-learn/rpc_test_server/pb"
)

const defaultUrl = ":9050"

type Client struct {
	conn    *grpc.ClientConn
	cert    credentials.TransportCredentials
	url     string
	traceId string
}

func NewRpcTestServerClient(traceId ...string) *Client {
	client := Client{}

	if traceId != nil {
		if len(traceId[0]) != 0 {
			client.traceId = traceId[0]
		}
	} else {
		client.traceId = uuid.NewV4().String()
	}

	var err error
	client.cert, err = credentials.NewClientTLSFromFile("C:\\Users\\alion\\GolandProjects\\rpc-learn\\rpc_test_server\\internal\\crypto\\server.csr", "")
	if err != nil {
		return nil
	}

	return &client
}

func (c *Client) Client() pb.TestServerClient {

	if c.conn != nil {
		if c.conn.GetState() == connectivity.Ready {
			c.ConnClose()
		}
	} else {
		c.conn = &grpc.ClientConn{}
	}

	if len(c.url) == 0 {
		c.SetUrl()
	}

	var err error
	c.conn, err = grpc.Dial(c.url, grpc.WithTransportCredentials(c.cert))
	if err != nil {
		return nil
	}

	return pb.NewTestServerClient(c.conn)
}

func (c *Client) ConnClose() {
	if c.conn != nil {
		_ = c.conn.Close()
	}
}

func (c *Client) SetUrl(url ...string) {
	c.url = defaultUrl

	if url != nil {
		if len(url[0]) != 0 {
			c.url = url[0]
		}
	}
}

func (c *Client) GetTraceId() string {
	return c.traceId
}
