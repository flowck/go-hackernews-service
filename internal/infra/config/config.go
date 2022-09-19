package config

import "github.com/kelseyhightower/envconfig"

type Cfg struct {
	Evironment           string `envconfig:"ENVIRONMENT"`
	DbDriverName         string `envconfig:"DB_DRIVER_NAME"`
	DbUrl                string `envconfig:"DB_URL"`
	DbMaxConnections     int    `envconfig:"DB_MAX_CONNECTIONS"`
	DbMaxIdleConnections int    `envconfig:"DB_MAX_IDLE_CONNECTIONS"`
}

func GetConfig() *Cfg {
	var cfg Cfg
	if err := envconfig.Process("", &cfg); err != nil {
		panic(err)
	}

	return &cfg
}
