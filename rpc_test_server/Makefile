dump:
	pg_dump \
		--create --schema-only \
		--file="${GOPATH}/src/rpc-learn/migration/01_test_db.sql" \
		--dbname="test" --host="localhost" --port="5432" --username="postgres"

xo:
	xo \
		pgsql://postgres:postgres@localhost/test?sslmode=disable \
		-int32-type int64 -uint32-type uint64 -f _test.xo.go \
		-o ./internal/models

proto:
	protoc \
		-I="${GOPATH}/rpc-learn/rpc_test_server" \
		--go_out=plugins=grpc:. \
		test_server.proto

.PHONY crypto:
	openssl genrsa \
		-out "${GOPATH}/rpc-learn/rpc_test_server/internal/crypto/server_rsa.key" 2048

	openssl req  \
		-new -x509 -sha256 \
		-key "${GOPATH}/rpc-learn/rpc_test_server/internal/crypto/server_rsa.key" \
		-out "${GOPATH}/rpc-learn/rpc_test_server/internal/crypto/server.csr"

	openssl x509 -x509toreq -days 365 \
		-in "${GOPATH}/rpc-learn/rpc_test_server/internal/crypto/server.csr" \
		-signkey "${GOPATH}/rpc-learn/rpc_test_server/internal/crypto/server_rsa.key" \
		 -out "${GOPATH}/rpc-learn/rpc_test_server/internal/crypto/server.req"

	openssl x509 \
		-req -days 365 \
		-in "${GOPATH}/rpc-learn/rpc_test_server/internal/crypto/server.req" \
		-signkey "${GOPATH}/rpc-learn/rpc_test_server/internal/crypto/server_rsa.key" \
		-out "${GOPATH}/rpc-learn/rpc_test_server/internal/crypto/server.crt"
