package pvz

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (s *serv) CloseReception(ctx context.Context, pvzId uuid.UUID) (*model.Reception, error) {
	lastReception, err := s.pvzRepository.GetLastReception(ctx, pvzId)
	if err != nil {
		return nil, err
	}

	if lastReception.Status == model.StatusClose {
		return nil, errors.New("reception is already closed")
	}

	err = s.pvzRepository.UpdateReception(ctx, lastReception.ID)
	if err != nil {
		return nil, err
	}
	lastReception.Status = model.StatusClose
	return lastReception, nil
}
