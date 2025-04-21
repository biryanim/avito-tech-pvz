package pvz

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (r *repo) GetLastReception(ctx context.Context, pvzId uuid.UUID) (*model.Reception, error) {
	builder := sq.Select(idColumnName, createdAtColumnName, pvzIdColumnName, statusColumnName).
		From(receptionsTableName).
		Where(sq.Eq{pvzIdColumnName: pvzId}).
		OrderBy(createdAtColumnName + " DESC").
		Limit(1).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var rec model.Reception
	err = r.db.DB().QueryRowContext(ctx, query, args...).Scan(&rec.ID, &rec.CreatedAt, &rec.PvzId, &rec.Status)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &rec, nil
}
