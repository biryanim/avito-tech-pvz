package access

import (
	"github.com/biryanim/avito-tech-pvz/internal/client/db"
	"github.com/biryanim/avito-tech-pvz/internal/repository"
)

const (
	accessesTableName = "accesses"

	idColumn              = "id"
	endpointAddressColumn = "endpoint_address"
	roleColumn            = "role"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.AccessRepository {
	return &repo{
		db: db,
	}
}
