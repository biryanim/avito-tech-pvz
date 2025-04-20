package pvz

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
	"time"
)

func (r *repo) CreateReception(ctx context.Context, pvzId uuid.UUID) (*model.Reception, error) {
	id := uuid.New()
	dateTime := time.Now().UTC()
	builder := sq.Insert(receptionsTableName).
		Columns(idColumnName, createdAtColumnName, pvzIdColumnName, statusColumnName).
		Values(id, dateTime, pvzId, model.StatusInProgress)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = r.pgx.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &model.Reception{
		ID:       id,
		PvzId:    pvzId,
		Status:   model.StatusInProgress,
		DateTime: dateTime,
	}, nil
}
