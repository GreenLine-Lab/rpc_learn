package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"rpc/rpc-test-server/api"
	"rpc/rpc-test-server/pb"
)

var dev bool
var protocol, port string

func init() {
	flag.BoolVar(&dev, "dev", false, "flag developer mode")
	flag.StringVar(&protocol, "protocol", "tcp", "listen server data transfer protocol")
	flag.StringVar(&port, "port", "9050", "listen server port")
}

func main() {
	flag.Parse()
	lis, err := net.Listen(protocol, net.JoinHostPort("", port))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	srv := api.NewTestServer()
	pb.RegisterTestServerServer(grpcServer, &srv)

	log.Println("Listen ... " + port + " port")
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func isDebug() bool {
	return dev || os.Getenv("DEBUG") == "true"
}
