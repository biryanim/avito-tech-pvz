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
	Create(ctx context.Context, pvz *model.Pvz) (uuid.UUID, error)
	CreateReception(ctx context.Context, pvzId uuid.UUID) (*model.Reception, error)
	GetLastReception(ctx context.Context, pvzId uuid.UUID) (*model.Reception, error)
	//AddProductToReception(ctx context.Context, reception *model.Reception) error

	//Get(ctx context.Context, uuid uuid.UUID) (*model.Pvz, error)
}
