package config

import (
	"errors"
	"os"
	"time"
)

const (
	jwtTokenSecretKey  = "JWT_SECRET_KEY"
	jwtTokenExpiration = "JWT_TOKEN_EXPIRATION"
)

type JWTConfig interface {
	TokenSecret() []byte
	TokenExpiration() time.Duration
}

type jwtConfig struct {
	tokenSecret []byte
	tokenExp    time.Duration
}

func NewJWTConfig() (JWTConfig, error) {
	tokenSecret := []byte(os.Getenv(jwtTokenSecretKey))
	if len(tokenSecret) == 0 {
		return nil, errors.New("missing JWT token secret")
	}

	tokenExp, err := time.ParseDuration(os.Getenv(jwtTokenExpiration))
	if err != nil {
		return nil, err
	}

	return &jwtConfig{
		tokenSecret: tokenSecret,
		tokenExp:    tokenExp,
	}, nil
}

func (j *jwtConfig) TokenSecret() []byte {
	return j.tokenSecret
}

func (j *jwtConfig) TokenExpiration() time.Duration {
	return j.tokenExp
}
