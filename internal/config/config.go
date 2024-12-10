package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	DbURL         string `env:"DB_URL"`
	ApiServerPort uint16 `env:"API_SERVER_PORT"`

	AuthSecretKey string `env:"AUTH_SECRET_KEY"`
	AuthTTL       int64  `env:"AUTH_TTL"`
}

func GetConfig() (Config, error) {
	return env.ParseAs[Config]()
}
