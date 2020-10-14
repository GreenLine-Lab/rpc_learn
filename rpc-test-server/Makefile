test-server:
	protoc -I="${GOPATH}/src/rpc/rpc-test-server/proto" \
		--go_out=plugins=grpc:. \
		test_server.proto