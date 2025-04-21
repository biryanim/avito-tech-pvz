package pvz

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
	"time"
)

func (r *repo) CreateProduct(ctx context.Context, product *model.ProductInfo) (*model.Product, error) {
	id := uuid.New()
	createdAt := time.Now().UTC()
	builder := sq.Insert(productsTableName).
		Columns(idColumnName, createdAtColumnName, receptionIdColumnName, typeColumnName).
		Values(id, createdAt, product.ReceptionId, product.Type).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &model.Product{
		ID:        id,
		Info:      *product,
		CreatedAt: createdAt,
	}, nil
}
