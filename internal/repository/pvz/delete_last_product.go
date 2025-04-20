package pvz

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (r *repo) DeleteLastProduct(ctx context.Context, receptionId uuid.UUID) error {
	builder := sq.Delete("products").
		Where(sq.Eq{receptionIdColumnName: receptionId}).
		OrderBy(createdAtColumnName + " DESC").
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pgx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}
