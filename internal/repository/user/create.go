package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/biryanim/avito-tech-pvz/internal/model"
)

func (r *repo) Create(ctx context.Context, user *model.UserRegistration) (string, error) {
	builder := sq.Insert(tableName).
		Columns(emailColumnName, passwordColumnName, roleColumnName).
		Values(user.Info.Email, user.Password, user.Info.Role).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	var id string
	err = r.pgx.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}
