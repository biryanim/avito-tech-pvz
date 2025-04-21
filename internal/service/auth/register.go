package auth

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/pkg/errors"
)

func (s *serv) Register(ctx context.Context, register *model.UserRegistration) (*model.User, error) {
	if !register.Info.Role.IsValid() {
		//TODO: добавить кастомные ошибки для моделек
		return nil, errors.New("invalid role")
	}

	hashedPassword, err := s.hashPassword(register.Password)
	if err != nil {
		return nil, err
	}

	register.Password = hashedPassword
	id, err := s.userRepository.Create(ctx, register)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID: id,
		Info: model.UserInfo{
			Email: register.Info.Email,
			Role:  register.Info.Role,
		},
	}, nil
}
