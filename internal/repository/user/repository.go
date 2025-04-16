package user

import (
	"github.com/biryanim/avito-tech-pvz/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	tableName = "users"

	idColumnName       = "id"
	emailColumnName    = "email"
	passwordColumnName = "password"
	roleColumnName     = "role"
)

type repo struct {
	pgx *pgxpool.Pool
}

func NewRepository(pgxPool *pgxpool.Pool) repository.UserRepository {
	return &repo{
		pgx: pgxPool,
	}
}
