package user

import (
	"github.com/biryanim/avito-tech-pvz/internal/client/db"
	"github.com/biryanim/avito-tech-pvz/internal/repository"
)

const (
	tableName = "users"

	idColumnName       = "id"
	emailColumnName    = "email"
	passwordColumnName = "password"
	roleColumnName     = "role"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{
		db: db,
	}
}
