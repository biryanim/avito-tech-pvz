package pvz

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (s *serv) CreateReception(ctx context.Context, pvzId uuid.UUID) (*model.Reception, error) {
	lastReception, err := s.pvzRepository.GetLastReception(ctx, pvzId)
	if err != nil {
		return nil, err
	}
	if lastReception.Status != model.StatusClose {
		return nil, errors.New("last reception is not closed")
	}

	reception, err := s.pvzRepository.CreateReception(ctx, pvzId)
	if err != nil {
		return nil, err
	}

	return reception, nil
}
