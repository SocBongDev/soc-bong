package config

import (
	"fmt"
	"log/slog"
	"net/url"
)

type DatabaseSecret struct {
	Url url.URL `env:"DB_URL" json:"db_url"`
}

var _ slog.LogValuer = (*DatabaseSecret)(nil)

func (d *DatabaseSecret) GetUrl() string {
	return d.Url.String()
}

func (d *DatabaseSecret) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("db_url", fmt.Sprintf("**%s**", d.Url.Host)),
	)
}
