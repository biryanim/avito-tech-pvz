package auth

import (
	"context"
	"errors"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/biryanim/avito-tech-pvz/internal/utils"
)

func (s *serv) Login(ctx context.Context, loginInfo *model.UserLoginInfo) (string, error) {
	user, err := s.userRepository.Get(ctx, loginInfo.Email)
	if err != nil {
		return "", err
	}

	if !utils.VerifyPassword(user.Password, loginInfo.Password) {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateToken(user.Info.Role, s.jwtConfig.TokenSecret(), s.jwtConfig.TokenExpiration())
	if err != nil {
		return "", err
	}
	return token, nil
}
