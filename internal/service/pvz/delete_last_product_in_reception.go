package pvz

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (s *serv) DeleteLastProductInReception(ctx context.Context, pvzId uuid.UUID) error {

	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		lastReception, err := s.pvzRepository.GetLastReception(ctx, pvzId)
		if lastReception == nil {
			return errors.New("no such pvz")
		}
		if err != nil {
			return err
		}

		if lastReception.Status != model.StatusInProgress {
			return errors.New("reception is already closed")
		}

		err = s.pvzRepository.DeleteLastProduct(ctx, lastReception.ID)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
