package repository

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.UserRegistration) (uuid.UUID, error)
	Get(ctx context.Context, email string) (*model.User, error)
}

type AccessRepository interface {
	GetList(ctx context.Context) ([]*model.AccessInfo, error)
}

type PvzRepository interface {
	Create(ctx context.Context, pvz *model.PVZInfo) (*model.PVZ, error)
	CreateReception(ctx context.Context, pvzId uuid.UUID) (*model.Reception, error)
	GetLastReception(ctx context.Context, pvzId uuid.UUID) (*model.Reception, error)
	CreateProduct(ctx context.Context, product *model.ProductInfo) (*model.Product, error)
	DeleteLastProduct(ctx context.Context, receptionId uuid.UUID) error
	UpdateReception(ctx context.Context, receptionId uuid.UUID) error
	GetListPVZ(ctx context.Context, pagination *model.Filter) ([]*model.PVZWithReceptions, error)
}
