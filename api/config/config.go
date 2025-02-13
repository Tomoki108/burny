package config

import "github.com/caarlos0/env/v11"

type Config struct {
	JwtSecret string `env:"JWT_SECRET"`
	Host      string `env:"HOST"`
}

var Conf Config

func Init() error {
	return env.Parse(&Conf)
}
