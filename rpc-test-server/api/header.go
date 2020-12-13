package api

import (
	"database/sql"
	"github.com/rs/zerolog"
	"rpc-learn/lib"
)

type TestServer struct {
	db     *sql.DB
	logger *zerolog.Logger
}

func NewTestServer(cfg *lib.EnvConfig) (*TestServer, error) {

	db, err := cfg.ConnectDB()
	if err != nil {
		return nil, err
	}

	return &TestServer{
		db: db,
	}, nil
}

func (srv *TestServer) DB() *sql.DB {
	return srv.db
}
