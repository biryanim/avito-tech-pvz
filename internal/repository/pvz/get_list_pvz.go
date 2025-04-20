package pvz

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
)

func (r *repo) GetListPVZ(ctx context.Context, pagination *model.Filter) ([]*model.PVZWithReceptions, error) {
	builder := sq.Select(
		pvzTableName+"."+idColumnName,
		pvzTableName+"."+cityColumnName,
		pvzTableName+"."+createdAtColumnName,
		receptionsTableName+"."+idColumnName,
		receptionsTableName+"."+pvzIdColumnName,
		receptionsTableName+"."+statusColumnName,
		receptionIdColumnName+"."+createdAtColumnName,
		productsTableName+"."+idColumnName,
		productsTableName+"."+receptionIdColumnName,
		productsTableName+"."+typeColumnName,
		productsTableName+"."+createdAtColumnName,
	).From(pvzTableName).
		LeftJoin(receptionsTableName+"ON"+pvzTableName+"."+idColumnName+"="+receptionsTableName+"."+pvzIdColumnName).
		LeftJoin(productsTableName+"ON"+receptionsTableName+"."+idColumnName+"="+productsTableName+"."+receptionIdColumnName).
		OrderBy(
			pvzTableName+"."+createdAtColumnName,
			receptionsTableName+"."+createdAtColumnName+" DESC",
			productsTableName+"."+createdAtColumnName+" DESC",
		).
		Limit(pagination.Limit).
		Offset((pagination.Page - 1) * pagination.Limit)

	if !pagination.StartDate.IsZero() {
		builder = builder.Where(sq.GtOrEq{receptionsTableName + "." + createdAtColumnName: pagination.StartDate})
	}
	if !pagination.EndDate.IsZero() {
		builder = builder.Where(sq.LtOrEq{receptionsTableName + "." + createdAtColumnName: pagination.EndDate})
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pgx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pvzMap := make(map[uuid.UUID]*model.PVZWithReceptions)
	for rows.Next() {
		var pvz model.PVZ
		var reception model.Reception
		var product model.Product

		err = rows.Scan(
			&pvz.ID, &pvz.Info.City, &pvz.RegistrationDate,
			&reception.ID, &reception.PvzId, &reception.Status, &reception.DateTime,
			&product.ID, &product.Info.ReceptionId, &product.Info.Type, &product.CreatedAt)
		if err != nil {
			return nil, err
		}

		if _, ok := pvzMap[pvz.ID]; !ok {
			pvzMap[pvz.ID] = &model.PVZWithReceptions{
				PVZ:        pvz,
				Receptions: []model.ReceptionsWithProducts{},
			}
		}

		foundReception := false
		for i := range pvzMap[pvz.ID].Receptions {
			if pvzMap[pvz.ID].Receptions[i].Reception.ID == reception.ID {
				pvzMap[pvz.ID].Receptions[i].Products = append(pvzMap[pvz.ID].Receptions[i].Products, product)
				foundReception = true
				break
			}
		}
		if !foundReception {
			receptionWithProducts := model.ReceptionsWithProducts{
				Reception: reception,
				Products:  []model.Product{},
			}
			receptionWithProducts.Products = append(receptionWithProducts.Products, product)
			pvzMap[pvz.ID].Receptions = append(pvzMap[pvz.ID].Receptions, receptionWithProducts)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	result := make([]*model.PVZWithReceptions, 0, len(pvzMap))
	for _, p := range pvzMap {
		result = append(result, p)
	}

	return result, nil
}
