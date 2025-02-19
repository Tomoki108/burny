package config

import "github.com/caarlos0/env/v11"

type Config struct {
	JwtSecret                   string `env:"JWT_SECRET"`
	Host                        string `env:"HOST"`
	Port                        string `env:"PORT"`
	DB_NAME                     string `env:"DB_NAME"`
	DB_USER                     string `env:"DB_USER"`
	DB_PASS                     string `env:"DB_PASS"`
	DB_INSTANCE_CONNECTION_NAME string `env:"DB_INSTANCE_CONNECTION_NAME"`
}

var Conf Config

func Init() error {
	return env.Parse(&Conf)
}
