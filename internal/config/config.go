package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPPort int `default:"8080" envconfig:"http_port"`
}

func New() (Config, error) {
	var c Config

	err := envconfig.Process("", &c)

	return c, err
}
