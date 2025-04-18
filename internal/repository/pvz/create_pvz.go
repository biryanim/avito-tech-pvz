package pvz

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

func (r *repo) Create(ctx context.Context, pvz *model.Pvz) (uuid.UUID, error) {
	if len(pvz.ID.String()) == 0 {
		pvz.ID = uuid.New()
	}

	if pvz.RegistrationDate.IsZero() {
		pvz.RegistrationDate = time.Now()
	}

	builder := sq.Insert(pvzTableName).
		Columns(idColumnName, registrationDateColumnName, cityColumnName).
		Values(pvz.ID, pvz.RegistrationDate, pvz.City).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "failed to build query")
	}

	_, err = r.pgx.Exec(ctx, query, args...)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "failed to execute query")
	}

	return pvz.ID, nil
}
