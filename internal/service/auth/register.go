package auth

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/pkg/errors"
)

func (s *serv) Register(ctx context.Context, register *model.UserRegistration) (*model.User, error) {
	if !register.Info.Role.IsValid() {
		return nil, model.ErrorInvalidRole
	}

	hashedPassword, err := s.hashPassword(register.Password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to hash password")
	}

	register.Password = hashedPassword
	id, err := s.userRepository.Create(ctx, register)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}

	return &model.User{
		ID: id,
		Info: model.UserInfo{
			Email: register.Info.Email,
			Role:  register.Info.Role,
		},
	}, nil
}
