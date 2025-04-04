package config

import (
	"errors"
	"os"
)

const (
	dsnEnvName = "PG_DSN"
)

type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
func NewPgConfig() (*pgConfig, error) {
	env := os.Getenv(dsnEnvName)
	if len(env) == 0 {
		return nil, errors.New("pg dsn not found")
	}
	return &pgConfig{dsn: env}, nil
}
