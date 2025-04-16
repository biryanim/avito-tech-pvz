package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/biryanim/avito-tech-pvz/internal/model"
)

func (r *repo) Get(ctx context.Context, email string) (*model.User, error) {
	builder := sq.Select(idColumnName, emailColumnName, passwordColumnName, roleColumnName).
		From(tableName).
		Where(sq.Eq{emailColumnName: email}).
		Limit(1).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var user model.User
	err = r.pgx.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Info.Email, &user.Password, &user.Info.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
