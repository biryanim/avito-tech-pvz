package pvz

import (
	"github.com/biryanim/avito-tech-pvz/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	pvzTableName        = "pvz"
	receptionsTableName = "receptions"
	productsTableName   = "products"

	idColumnName               = "id"
	registrationDateColumnName = "registration_date"
	cityColumnName             = "city"
	pvzIdColumnName            = "pvz_id"
	receptionIdColumnName      = "reception_id"
	dateTimeColumnName         = "date_time"
	statusColumnName           = "status"
)

type repo struct {
	pgx *pgxpool.Pool
}

func NewRepository(pgx *pgxpool.Pool) repository.PvzRepository {
	return &repo{
		pgx: pgx,
	}
}
