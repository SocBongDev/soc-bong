package config

type DatabaseSecret struct {
	Url string `env:"DB_URL"`
}
