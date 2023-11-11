package apiserver

import "rest-api/internal/app/nlab"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Nlab     *nlab.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Nlab:     nlab.NewConfig(),
	}
}
