package config

import (
	"github.com/blocktop/mp-common/config"
	"github.com/caarlos0/env"
)

type Config struct {
	config.BaseConfig
	StaticPath      string `env:"MP_WEB_STATIC_PATH"`
	StellarTOMLPath string `env:"MP_WEB_STELLAR_TOML_PATH"`
}

var cfg *Config

func init() {
	cfg = &Config{}
	cfg.Parse()
}

func GetConfig() *Config {
	return cfg
}

func (c *Config) Parse() {
	c.BaseConfig.Parse()

	err := env.Parse(c)
	if err != nil {
		panic(err)
	}
}
