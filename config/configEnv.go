package config

import "github.com/kelseyhightower/envconfig"

type ConfigEnv struct {
	Env struct {
		User string `envconfig:"USER"`
		Home string `envconfig:"HOME"`
	}
}

func ConfigHelperEnv() ConfigEnv {
	var cfg ConfigEnv
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
