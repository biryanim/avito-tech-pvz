package pvz

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (s *serv) DeleteLastProductInReception(ctx context.Context, pvzId uuid.UUID) error {
	lastReception, err := s.pvzRepository.GetLastReception(ctx, pvzId)
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
}
