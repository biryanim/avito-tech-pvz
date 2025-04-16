package auth

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/biryanim/avito-tech-pvz/internal/utils"
)

func (s *serv) Register(ctx context.Context, registerInfo *model.UserRegistration) (*model.User, error) {
	hashedPassword, err := utils.HashPassword(registerInfo.Password)
	if err != nil {
		return nil, err
	}

	registerInfo.Password = hashedPassword
	id, err := s.userRepository.Create(ctx, registerInfo)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID: id,
		Info: model.UserInfo{
			Email: registerInfo.Info.Email,
			Role:  registerInfo.Info.Role,
		},
	}, nil
}
