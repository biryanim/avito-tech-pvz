package pvz

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
)

func (r *repo) GetLastReception(ctx context.Context, pvzId uuid.UUID) (*model.Reception, error) {
	builder := sq.Select(idColumnName, dateTimeColumnName, pvzIdColumnName, statusColumnName).
		From(receptionsTableName).
		Where(sq.Eq{pvzIdColumnName: pvzId}).
		OrderBy(dateTimeColumnName + " DESC").
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var rec model.Reception
	err = r.pgx.QueryRow(ctx, query, args...).Scan(&rec.ID, &rec.DateTime, &rec.PvzId, &rec.Status)
	if err != nil {
		return nil, err
	}

	return &rec, nil
}
