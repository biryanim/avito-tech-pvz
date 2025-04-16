package auth

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/biryanim/avito-tech-pvz/internal/utils"
)

func (s *serv) DummyLogin(ctx context.Context, role string) (string, error) {
	token, err := utils.GenerateToken(model.UserInfo{Email: "", Role: role}, s.jwtConfig.TokenSecret(), s.jwtConfig.TokenExpiration())
	if err != nil {
		return "", err
	}
	return token, nil
}
