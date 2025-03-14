package config

import "github.com/caarlos0/env/v11"

// NOTE: envDefault is used for local environment and ci workflow environment.
type Config struct {
	JwtSecret string `env:"JWT_SECRET" envDefault:"random-str"`
	Host      string `env:"HOST" envDefault:"localhost"`
	Port      string `env:"PORT" envDefault:"1323"`
	DB_HOST   string `env:"DB_HOST" envDefault:"localhost"`
	DB_NAME   string `env:"DB_NAME" envDefault:"burny_db"`
	DB_USER   string `env:"DB_USER" envDefault:"burny_user"`
	DB_PASS   string `env:"DB_PASSWORD" envDefault:"pass"`
}

var Conf Config

func Init() error {
	return env.Parse(&Conf)
}
