package auth

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/pkg/errors"
)

func (s *serv) DummyLogin(ctx context.Context, role string) (string, error) {

	userRole := model.Role(role)
	if !userRole.IsValid() {
		//TODO: добавить кастомные ошибки для моделек
		return "", errors.New("invalid role")
	}
	token, err := s.generateToken(userRole, s.jwtConfig.TokenSecret(), s.jwtConfig.TokenExpiration())
	if err != nil {
		return "", err
	}
	return token, nil
}
