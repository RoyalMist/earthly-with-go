package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

const (
	ApiHost   = "API_HOST"
	dbDriver  = "DB_DRIVER"
	dbUrl     = "DB_URL"
	JwtSecret = "JWT_SECRET"
)

type Config struct {
	apiHost   string
	dbDriver  string
	dbUrl     string
	jwtSecret string
}

func (c Config) JwtSecret() string {
	return c.jwtSecret
}

func (c Config) DbUrl() string {
	return c.dbUrl
}

func (c Config) DbDriver() string {
	return c.dbDriver
}

func (c Config) ApiHost() string {
	return c.apiHost
}

// Module makes the injectable available for FX.
var Module = fx.Provide(New)

// New creates a new injectable.
func New() *Config {
	viper.SetDefault(ApiHost, ":4000")
	viper.SetDefault(dbDriver, "postgres")
	viper.SetDefault(dbUrl, "host=localhost port=5432 user=postgres dbname=change password=postgres sslmode=disable")
	viper.SetDefault(JwtSecret, "change")
	viper.AutomaticEnv()

	return &Config{
		apiHost:   viper.GetString(ApiHost),
		dbDriver:  viper.GetString(dbDriver),
		dbUrl:     viper.GetString(dbUrl),
		jwtSecret: viper.GetString(JwtSecret),
	}
}
