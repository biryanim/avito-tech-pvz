package auth

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/pkg/errors"
)

func (s *serv) DummyLogin(ctx context.Context, role model.Role) (string, error) {
	token, err := s.generateToken(role, s.jwtConfig.TokenSecret(), s.jwtConfig.TokenExpiration())
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return token, nil
}
