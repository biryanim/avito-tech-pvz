package pvz

import (
	"github.com/biryanim/avito-tech-pvz/internal/client/db"
	"github.com/biryanim/avito-tech-pvz/internal/repository"
)

const (
	pvzTableName        = "pvz"
	receptionsTableName = "receptions"
	productsTableName   = "products"

	idColumnName          = "id"
	cityColumnName        = "city"
	pvzIdColumnName       = "pvz_id"
	receptionIdColumnName = "reception_id"
	statusColumnName      = "status"
	typeColumnName        = "type"
	createdAtColumnName   = "created_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.PvzRepository {
	return &repo{
		db: db,
	}
}
