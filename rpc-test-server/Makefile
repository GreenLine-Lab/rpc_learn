dump:
	pg_dump \
		--create --schema-only \
		--file="${GOPATH}/src/rpc/migration/01_test_db.sql" \
		--dbname="test" --host="localhost" --port="5432" --username="postgres" --password="postgres"

test-server:
	protoc \
		-I="${GOPATH}/src/rpc/rpc-test-server/proto" \
		--go_out=plugins=grpc:. \
		test_server.proto