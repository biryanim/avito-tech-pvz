package pvz

import (
	"context"
	"database/sql"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (s *serv) DeleteLastProductInReception(ctx context.Context, pvzId uuid.UUID) error {

	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		lastReception, err := s.pvzRepository.GetLastReception(ctx, pvzId)
		if lastReception == nil {
			return model.ErrNoSuchPvz
		}
		if err != nil {
			return err
		}

		if lastReception.Status != model.StatusInProgress {
			return model.ErrNoOpenReceptions
		}

		err = s.pvzRepository.DeleteLastProduct(ctx, lastReception.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return model.ErrNoProducts
			}
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
