package converter

import (
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/model"
)

func ToUserRegistrationModelFromRegistrationDTO(registerInfo dto.RegisterRequest) *model.UserRegistration {
	return &model.UserRegistration{
		Info: model.UserInfo{
			Email: registerInfo.Email,
			Role:  registerInfo.Role,
		},
		Password: registerInfo.Password,
	}
}

func ToRegistrationRespFromUserModel(user *model.User) *dto.RegisterResponse {
	return &dto.RegisterResponse{
		ID:    user.ID.String(),
		Email: user.Info.Email,
		Role:  user.Info.Role,
	}
}

func ToLoginInfoFromDTO(userInfo *dto.LoginRequest) *model.UserLoginInfo {
	return &model.UserLoginInfo{
		Email:    userInfo.Email,
		Password: userInfo.Password,
	}
}
