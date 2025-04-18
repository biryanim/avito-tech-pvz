package auth

import (
	"github.com/biryanim/avito-tech-pvz/internal/config"
	"github.com/biryanim/avito-tech-pvz/internal/repository"
	"github.com/biryanim/avito-tech-pvz/internal/service"
)

type serv struct {
	userRepository   repository.UserRepository
	accessRepository repository.AccessRepository
	jwtConfig        config.JWTConfig
}

func NewService(userRepository repository.UserRepository, accessRepository repository.AccessRepository, jwtCfg config.JWTConfig) service.AuthService {
	return &serv{
		userRepository:   userRepository,
		accessRepository: accessRepository,
		jwtConfig:        jwtCfg,
	}
}
