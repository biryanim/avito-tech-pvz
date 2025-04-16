package service

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
)

type AuthService interface {
	Register(ctx context.Context, registerInfo *model.UserRegistration) (*model.User, error)
	Login(ctx context.Context, loginInfo *model.UserLoginInfo) (string, error)
	DummyLogin(ctx context.Context, role string) (string, error)
}
