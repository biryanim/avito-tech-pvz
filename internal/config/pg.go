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

func NewPGConfig() (PGConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("pg DSN not found")
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (p *pgConfig) DSN() string {
	return p.dsn
}
