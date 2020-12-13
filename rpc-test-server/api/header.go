package api

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"rpc-learn/lib"
	"rpc-learn/lib/zlog"
)

type TestServer struct {
	db     *sql.DB
	logger zerolog.Logger
}

func NewTestServer(cfg *lib.EnvConfig) (*TestServer, error) {

	db, err := cfg.ConnectDB()
	if err != nil {
		return nil, err
	}

	logger := zlog.ZLogger{
		NoColor:    !cfg.ServiceDevMod,
		JSONFormat: false,
	}

	return &TestServer{
		db:     db,
		logger: logger.Init(),
	}, nil
}

func (srv *TestServer) DB() *sql.DB {
	return srv.db
}

func (srv *TestServer) GetLogger() zerolog.Logger {
	return srv.logger
}
