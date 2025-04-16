package config

import (
	"errors"
	"os"
)

const (
	dsnEnvName = "PG_DSN"
)

type PGConfig struct {
	dsn string
}

func NewPGConfig() (*PGConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("pg DSN not found")
	}

	return &PGConfig{
		dsn: dsn,
	}, nil
}

func (p *PGConfig) DSN() string {
	return p.dsn
}
