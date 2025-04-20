package service

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
)

type AuthService interface {
	Register(ctx context.Context, registerInfo *model.UserRegistration) (*model.User, error)
	Login(ctx context.Context, loginInfo *model.UserLoginInfo) (string, error)
	DummyLogin(ctx context.Context, role string) (string, error)
	Check(ctx context.Context, token, method, endpointAddress string) (bool, error)
}

type PVZService interface {
	CreatePVZ(ctx context.Context, pvz *model.PVZInfo) (*model.PVZ, error)
	CreateReception(ctx context.Context, pvzId uuid.UUID) (*model.Reception, error)
	AddProductToReception(ctx context.Context, productPVZ *model.ProductPVZ) (*model.Product, error)
	CloseReception(ctx context.Context, pvzId uuid.UUID) (*model.Reception, error)
	DeleteLastProductInReception(ctx context.Context, pvzId uuid.UUID) error
	GetListPVZs(ctx context.Context, pagination *model.Filter) ([]*model.PVZWithReceptions, error)
}
