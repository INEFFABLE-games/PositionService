package config

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

// Config structure for parse environment variables
type Config struct {
	UsersPort   string `env:"USERSPORT,required,notEmpty"`
	GrpcPort    string `env:"GRPCPORT,required,notEmpty"`
	PosGrpcPort string `env:"POSGRPCPORT,required,notEmpty"`
	PsqlURI     string `env:"PSQL_URI,required,notEmpty"`
}

// NewConfig create new config object.
func NewConfig() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.WithFields(log.Fields{
			"handler": "config",
			"action":  "initialize",
		}).Errorf("unable to pars environment variables %v,", err)
	}

	return &cfg
}
