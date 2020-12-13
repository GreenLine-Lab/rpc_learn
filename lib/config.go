package lib

import (
	"database/sql"
	"fmt"
)

type EnvConfig struct {
	ServiceName   string `env:"NAME"`
	ServiceDevMod bool   `env:"DEV_MOD" envDefault:"false"`

	ServiceHost string `env:"HOST" envDefault:"localhost"`
	ServicePort string `env:"PORT"`

	SqlType     string `env:"SQL_TYPE" envDefault:"postgres"`
	SqlHost     string `env:"SQL_HOST" envDefault:"localhost"`
	SqlPort     string `env:"SQL_PORT" envDefault:"5432"`
	SqlUser     string `env:"SQL_USER" envDefault:"postgres"`
	SqlPassword string `env:"SQL_PASSWORD" envDefault:"postgres"`
	SqlDatabase string `env:"SQL_DATABASE"`
}

func (cfg *EnvConfig) ConnectDB() (*sql.DB, error) {
	conn, err := sql.Open(cfg.SqlType,
		fmt.Sprintf(""+
			"host=%s "+
			"port=%s "+
			"user=%s "+
			"password=%s "+
			"dbname=%s sslmode=disable",
			cfg.SqlHost, cfg.SqlPort, cfg.SqlUser, cfg.SqlPassword, cfg.SqlDatabase))

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
