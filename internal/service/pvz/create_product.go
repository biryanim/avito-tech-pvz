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

	var (
		lastReception *model.Reception
		product       *model.Product
	)

	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var err error
		lastReception, err = s.pvzRepository.GetLastReception(ctx, productPVZ.PvzId)
		if err != nil {
			return err
		}

		if lastReception.Status != model.StatusInProgress {
			return model.ErrNoOpenReceptions
		}

		productInfo := &model.ProductInfo{
			Type:        productPVZ.Type,
			ReceptionId: lastReception.ID,
		}

		product, err = s.pvzRepository.CreateProduct(ctx, productInfo)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return product, nil
}
