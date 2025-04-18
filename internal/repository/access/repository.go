package access

import (
	"github.com/biryanim/avito-tech-pvz/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	accessesTableName = "accesses"

	idColumn              = "id"
	endpointAddressColumn = "endpoint_address"
	roleColumn            = "role"
)

type repo struct {
	pgx *pgxpool.Pool
}

func NewRepository(pgxPool *pgxpool.Pool) repository.AccessRepository {
	return &repo{
		pgx: pgxPool,
	}
}
