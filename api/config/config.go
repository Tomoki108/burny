package config

import "github.com/caarlos0/env/v11"

type Config struct {
	JwtSecret string `env:"JWT_SECRET"`
}

var Conf Config

func Init() error {
	return env.Parse(&Conf)
}
