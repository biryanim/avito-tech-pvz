package repository

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
)

//TODO: посмотреть возможно вместо string нужно использовать UUID

type UserRepository interface {
	Create(ctx context.Context, user *model.UserRegistration) (uuid.UUID, error)
	Get(ctx context.Context, email string) (*model.User, error)
}
