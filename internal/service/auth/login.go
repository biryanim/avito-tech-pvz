package auth

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/pkg/errors"
)

func (s *serv) Login(ctx context.Context, loginInfo *model.UserLoginInfo) (string, error) {
	user, err := s.userRepository.Get(ctx, loginInfo.Email)
	if err != nil {
		return "", errors.Wrap(err, model.ErrorUserNotFound.Error())
	}

	if !s.verifyPassword(user.Password, loginInfo.Password) {
		return "", model.ErrorInvalidPassword
	}

	token, err := s.generateToken(user.Info.Role, s.jwtConfig.TokenSecret(), s.jwtConfig.TokenExpiration())
	if err != nil {
		return "", errors.Wrap(err, "failed to generate token")
	}
	return token, nil
}
