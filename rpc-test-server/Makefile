dump:
	pg_dump \
		--create --schema-only \
		--file="${GOPATH}/src/rpc/migration/01_test_db.sql" \
		--dbname="test" --host="localhost" --port="5432" --username="postgres"

xo:
	xo \
		pgsql://postgres:postgres@localhost/test \
		-int32-type int64 -uint32-type uint64 \
		-o ${GOPATH}/src/rpc/rpc-test-server/models

grpc-test-server:
	protoc \
		-I="${GOPATH}/src/rpc/rpc-test-server/proto" \
		--go_out=plugins=grpc:. \
		test_server.proto