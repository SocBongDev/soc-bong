package config

import (
	"fmt"
	"log"
	"log/slog"
	"strings"

	"github.com/SocBongDev/soc-bong/internal/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

func generateMaskedString(input string) string {
	length := len(input)
	masked := strings.Repeat("*", length)
	return masked
}

const (
	DEV  = "dev"
	PROD = "prod"
)

var (
	_      slog.LogValuer = (*Config)(nil)
	config *Config
)

type Config struct {
	Auth0Secret
	DatabaseSecret
	ClientOriginUrl string `env:"CLIENT_ORIGIN_URL" env-default:"http://localhost:5173" json:"client_origin_url"`
	Env             string `env:"ENV" env-default:"dev" json:"env"`
	Port            int    `env:"PORT" env-default:"5000" json:"port"`
}

func (c *Config) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("port", c.Port),

		slog.String("auth0_audience", generateMaskedString(c.Audience)),
		slog.String("auth0_domain", generateMaskedString(c.Domain)),
		slog.String("client_origin_url", c.ClientOriginUrl),
		slog.String("db_url", fmt.Sprintf("**%s**", c.Url.Host)),
		slog.String("env", c.Env),
	)
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

func init() {
	var err error
	config, err = loadConfig()
	if err != nil {
		logger.Error("Failed to load config: ", err)
	}

	logger.Info("Init config success", "config", config)
}

// GetConfig returns the singleton instance of Config
func GetConfig() *Config {
	return config
}

func loadConfig() (*Config, error) {
	config := &Config{}
	if err := cleanenv.ReadEnv(config); err != nil {
		return nil, err
	}

	if config.Env == "dev" {
		if err := cleanenv.ReadConfig(".env", config); err != nil {
			return nil, err
		}
	}
	return config, nil
}
