package pvz

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (r *repo) DeleteLastProduct(ctx context.Context, receptionId uuid.UUID) error {
	subquery := sq.Select(idColumnName).
		From(productsTableName).
		Where(sq.Eq{receptionIdColumnName: receptionId}).
		OrderBy(createdAtColumnName + " DESC").
		Limit(1).PlaceholderFormat(sq.Dollar)

	subSql, subArgs, err := subquery.ToSql()
	if err != nil {
		return err
	}

	builder := sq.Delete(productsTableName).
		Where(idColumnName+" = ("+subSql+")", subArgs...).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	tag, err := r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return sql.ErrNoRows
	}
	return nil
}
