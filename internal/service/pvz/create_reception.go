package pvz

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
)

func (s *serv) CreateReception(ctx context.Context, pvzId uuid.UUID) (*model.Reception, error) {
	var (
		lastReception *model.Reception
		newReception  *model.Reception
	)
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var err error
		lastReception, err = s.pvzRepository.GetLastReception(ctx, pvzId)
		if err != nil {
			return err
		}

		if lastReception != nil && lastReception.Status != model.StatusClose {
			return model.ErrLastReceptionNotClosed
		}

		newReception, err = s.pvzRepository.CreateReception(ctx, pvzId)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return newReception, nil
}
