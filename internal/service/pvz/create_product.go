package pvz

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/pkg/errors"
)

func (s *serv) AddProductToReception(ctx context.Context, productPVZ *model.ProductPVZ) (*model.Product, error) {
	if !productPVZ.Type.IsValid() {
		// TODO: добавить кастомные ошибки при валидации полей бизнес моделек
		return nil, errors.New("invalid product type")
	}
	lastReception, err := s.pvzRepository.GetLastReception(ctx, productPVZ.PvzId)
	if err != nil {
		return nil, err
	}

	if lastReception.Status != model.StatusInProgress {
		return nil, errors.New("reception is already closed")
	}

	productInfo := &model.ProductInfo{
		Type:        productPVZ.Type,
		ReceptionId: lastReception.ID,
	}

	product, err := s.pvzRepository.CreateProduct(ctx, productInfo)
	if err != nil {
		return nil, err
	}

	return product, nil
}
