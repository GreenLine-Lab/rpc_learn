package main

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"google.golang.org/grpc"
	"net"
	"rpc-learn/lib"
	"rpc-learn/lib/zlog"
	"rpc-learn/rpc-test-server/api"
	"rpc-learn/rpc-test-server/pb"
)

func main() {
	flag.Parse()

	cfg := lib.EnvConfig{
		ServiceName: "TestServer",
		ServicePort: "9050",
		SqlDatabase: "test",
	}

	if err := env.Parse(&cfg); err != nil {
		panic(err.Error())
	}

	srv, err := api.NewTestServer(&cfg)
	if err != nil {
		panic(err.Error())
	}

	log := zlog.SrvLogger(srv.GetLogger(), cfg.ServiceName)

	defer func() {
		if err := srv.DB().Close(); err != nil {
			log.Error().Msgf("Unable close database connection: %s", err.Error())
		}
	}()

	lis, err := net.Listen("tcp", net.JoinHostPort(cfg.ServiceHost, cfg.ServicePort))
	if err != nil {
		log.Fatal().Msgf("Unable create new listener: %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTestServerServer(grpcServer, srv)

	log.Info().Msgf("Listen %s ...", net.JoinHostPort(cfg.ServiceHost, cfg.ServicePort))
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal().Msgf("Unable serve gRPC server: %s", err.Error())
	}
}
