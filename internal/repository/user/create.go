package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
)

func (r *repo) Create(ctx context.Context, user *model.UserRegistration) (uuid.UUID, error) {
	id := uuid.New()

	builder := sq.Insert(tableName).
		Columns(idColumnName, emailColumnName, passwordColumnName, roleColumnName).
		Values(id, user.Info.Email, user.Password, user.Info.Role).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return uuid.Nil, err
	}

	_, err = r.pgx.Exec(ctx, query, args...)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
