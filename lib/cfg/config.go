package cfg

import "sync"

type Config struct {
	Service string `env:"-"`

	// Orchestrator connection settings
	OrcHost           string `env:"ORC_HOST" envDefault:"192.168.20.114"`
	OrcPort           string `env:"ORC_PORT" envDefault:"9010"`
	OrcTlsSrvOverride string `env:"ORC_TLS_OVERRIDE" envDefault:""`

	// Listener settings
	LisHost string `env:"LISTEN_HOST"`
	LisPort string `env:"LISTEN_PORT"`

	// Db settings
	SqlType string `env:"SQL_TYPE" envDefault:"postgres"`
	SqlHost string `env:"SQL_HOST" envDefault:"localhost"`
	SqlPort string `env:"SQL_PORT" envDefault:"5432"`
	SqlUser string `env:"SQL_USER" envDefault:"postgres"`
	SqlPass string `env:"SQL_PASS" envDefault:"postgres"`
	SqlDb   string `env:"SQL_DB"`

	// Develop settings
	IsDev    bool `env:"IS_DEV" envDefault:"false"`
	IsDocker bool `env:"IS_DOCKER" envDefault:"false"`

	// Other settings
	Internal map[string]string `env:"-"`
	mu       sync.Mutex        `env:"-"`
}
