package auth

import (
	"github.com/biryanim/avito-tech-pvz/internal/config"
	"github.com/biryanim/avito-tech-pvz/internal/repository"
	"github.com/biryanim/avito-tech-pvz/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
	jwtConfig      config.JWTConfig
}

func NewService(userRepository repository.UserRepository, jwtCfg config.JWTConfig) service.AuthService {
	return &serv{
		userRepository: userRepository,
		jwtConfig:      jwtCfg,
	}
}
