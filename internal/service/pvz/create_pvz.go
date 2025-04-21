package pvz

import (
	"context"

	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/pkg/errors"
)

func (s *serv) CreatePVZ(ctx context.Context, pvzInfo *model.PVZInfo) (*model.PVZ, error) {
	pvz, err := s.pvzRepository.Create(ctx, pvzInfo)
	if err != nil {
		return nil, errors.Wrap(err, "create pvz")
	}

	return pvz, nil
}
