package pvz

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

func (r *repo) Create(ctx context.Context, pvzInfo *model.PVZInfo) (*model.PVZ, error) {
	id := uuid.New()
	createdAt := time.Now().UTC()

	builder := sq.Insert(pvzTableName).
		Columns(idColumnName, createdAtColumnName, cityColumnName).
		Values(id, createdAt, pvzInfo.City).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	_, err = r.pgx.Exec(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return &model.PVZ{
		ID:               id,
		RegistrationDate: createdAt,
		Info:             *pvzInfo,
	}, nil
}
