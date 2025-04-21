package auth

import (
	"context"
	"errors"
	"github.com/biryanim/avito-tech-pvz/internal/model"
)

func (s *serv) Login(ctx context.Context, loginInfo *model.UserLoginInfo) (string, error) {
	user, err := s.userRepository.Get(ctx, loginInfo.Email)
	if err != nil {
		return "", err
	}

	if !s.verifyPassword(user.Password, loginInfo.Password) {
		return "", errors.New("invalid password")
	}

	token, err := s.generateToken(user.Info.Role, s.jwtConfig.TokenSecret(), s.jwtConfig.TokenExpiration())
	if err != nil {
		return "", err
	}
	return token, nil
}
