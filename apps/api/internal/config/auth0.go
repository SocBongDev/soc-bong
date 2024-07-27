package config

import "log/slog"

var _ slog.LogValuer = (*Auth0Secret)(nil)

type Auth0Secret struct {
	Audience string `env:"AUTH0_AUDIENCE" json:"auth0_audience"`
	Domain   string `env:"AUTH0_DOMAIN" json:"auth0_domain"`
}

func (a *Auth0Secret) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("auth0_audience", generateMaskedString(a.Audience)),
		slog.String("auth0_domain", generateMaskedString(a.Domain)),
	)
}
