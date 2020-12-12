package lib

type EnvConfig struct {
	ServiceName   string
	ServiceDevMod bool `env:"DEV_MOD" envDefault:"false"`

	ServiceHost string `env:"HOST" envDefault:"localhost"`
	ServicePort string `env:"PORT"`

	SqlType     string `env:"SQL_TYPE" envDefault:"postgres"`
	SqlHost     string `env:"SQL_HOST" envDefault:"localhost"`
	SqlPort     string `env:"SQL_PORT" envDefault:"5432"`
	SqlUser     string `env:"SQL_USER" envDefault:"postgres"`
	SqlPassword string `env:"SQL_PASSWORD" envDefault:"postgres"`
}
