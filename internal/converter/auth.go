package converter

import (
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/pkg/errors"
)

func ToUserRegistrationModelFromRegistrationDTO(registerInfo *dto.RegisterRequest) *model.UserRegistration {
	return &model.UserRegistration{
		Info: model.UserInfo{
			Email: registerInfo.Email,
			Role:  model.Role(registerInfo.Role),
		},
		Password: registerInfo.Password,
	}
}

func ToRegistrationRespFromUserModel(user *model.User) *dto.RegisterResponse {
	return &dto.RegisterResponse{
		ID:    user.ID.String(),
		Email: user.Info.Email,
		Role:  string(user.Info.Role),
	}
}

func ToLoginInfoFromDTO(userInfo *dto.LoginRequest) *model.UserLoginInfo {
	return &model.UserLoginInfo{
		Email:    userInfo.Email,
		Password: userInfo.Password,
	}
}

func ToRoleModel(role string) (model.Role, error) {
	var r model.Role
	r = model.Role(role)
	if !r.IsValid() {
		return "", errors.New("invalid role")
	}

	return r, nil
}
