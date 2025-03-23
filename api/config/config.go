package config

import "github.com/caarlos0/env/v11"

// NOTE: envDefault is used for local environment and ci workflow environment.
type Config struct {
	JwtSecret  string `env:"JWT_SECRET" envDefault:"random-str"`
	Host       string `env:"HOST" envDefault:"localhost"`
	Port       string `env:"PORT" envDefault:"1323"`
	DB_HOST    string `env:"DB_HOST" envDefault:"localhost"`
	DB_NAME    string `env:"DB_NAME" envDefault:"burny_db"`
	DB_USER    string `env:"DB_USER" envDefault:"burny_user"`
	DB_PASS    string `env:"DB_PASSWORD" envDefault:"pass"`
	AWS_REGION string `env:"AWS_REGION" envDefault:"ap-northeast-1"` // for AWS SES
}

func (c Config) APIBaseURL() string {
	if c.Host == "localhost" {
		return "http://" + c.Host + ":" + c.Port + "/api/v1"
	}
	return "https://" + c.Host + "/api/v1"
}

var Conf Config

func Init() error {
	return env.Parse(&Conf)
}
