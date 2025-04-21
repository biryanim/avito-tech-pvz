package pvz

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
)

func (r *repo) UpdateReception(ctx context.Context, receptionId uuid.UUID) error {
	builder := sq.Update(receptionsTableName).
		Set(statusColumnName, model.StatusClose).
		Where(sq.Eq{idColumnName: receptionId}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}
