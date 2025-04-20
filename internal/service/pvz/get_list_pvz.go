package pvz

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
)

func (s *serv) GetListPVZs(ctx context.Context, pagination *model.Filter) ([]*model.PVZWithReceptions, error) {
	res, err := s.pvzRepository.GetListPVZ(ctx, pagination)
	if err != nil {
		return nil, err
	}
	return res, nil
}
