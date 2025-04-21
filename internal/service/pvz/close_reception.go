package pvz

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (s *serv) CloseReception(ctx context.Context, pvzId uuid.UUID) (*model.Reception, error) {
	var (
		lastReception *model.Reception
	)

	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var err error
		lastReception, err = s.pvzRepository.GetLastReception(ctx, pvzId)
		if err != nil {
			return err
		}

		if lastReception.Status == model.StatusClose {
			return errors.New("reception is already closed")
		}

		err = s.pvzRepository.UpdateReception(ctx, lastReception.ID)
		if err != nil {
			return err
		}
		lastReception.Status = model.StatusClose
		return nil
	})

	if err != nil {
		return nil, err
	}

	return lastReception, nil
}
