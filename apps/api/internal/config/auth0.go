package config

type Auth0Secret struct {
	Audience string `env:"AUTH0_AUDIENCE"`
	Domain   string `env:"AUTH0_DOMAIN"`
	ClientId string `env:"AUTH0_CLIENT_ID"`
	Secret   string `env:"AUTH0_CLIENT_SECRET"`
}
