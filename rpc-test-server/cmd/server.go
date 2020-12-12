package main

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"google.golang.org/grpc"
	"log"
	"net"
	"rpc-learn/lib"
	"rpc-learn/rpc-test-server/api"
	"rpc-learn/rpc-test-server/pb"
)

func main() {
	flag.Parse()

	cfg := lib.EnvConfig{
		ServicePort: "9050",
	}
	if err := env.Parse(&cfg); err != nil {
		panic(err.Error())
	}

	lis, err := net.Listen("tcp", net.JoinHostPort(cfg.ServiceHost, cfg.ServicePort))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	srv := api.NewTestServer()
	pb.RegisterTestServerServer(grpcServer, &srv)

	log.Println("Listen ... " + cfg.ServicePort + " port")
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
