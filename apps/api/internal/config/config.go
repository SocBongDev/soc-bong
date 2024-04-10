package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Auth0Secret
	DatabaseSecret
	ClientOriginUrl string `env:"CLIENT_ORIGIN_URL" env-default:"http://localhost:5173"`
	Env             string `env:"ENV"               env-default:"dev"`
	Port            int    `env:"PORT"              env-default:"5000"`
}

func New() (*Config, error) {
	config := &Config{}
	if err := cleanenv.ReadEnv(config); err != nil {
		log.Fatal("ReadEnv failed: ", err)
		return nil, err
	}

	if config.Env == "dev" {
		if err := cleanenv.ReadConfig(".env", config); err != nil {
			log.Fatal("ReadConfig from .env failed: ", err)
			return nil, err
		}
	}

	return config, nil
}
