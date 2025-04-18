package access

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/biryanim/avito-tech-pvz/internal/model"
)

func (r *repo) GetList(ctx context.Context) ([]*model.AccessInfo, error) {
	builder := sq.Select(idColumn, endpointAddressColumn, roleColumn).
		PlaceholderFormat(sq.Dollar).
		From(accessesTableName)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pgx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*model.AccessInfo

	for rows.Next() {
		var item model.AccessInfo
		err = rows.Scan(&item.Id, &item.EndpointAddress, &item.Role)
		if err != nil {
			return nil, err
		}
		result = append(result, &item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
