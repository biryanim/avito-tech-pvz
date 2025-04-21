package converter

import (
	"github.com/biryanim/avito-tech-pvz/internal/api/dto"
	"github.com/biryanim/avito-tech-pvz/internal/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

func ToPVZInfoFromDTO(pvzCreateDTO *dto.PVZCreateRequest) *model.PVZInfo {
	return &model.PVZInfo{
		City: model.City(pvzCreateDTO.City),
	}
}

func ToPVZResponseFromPVZ(pvz *model.PVZ) *dto.PVZResponse {
	return &dto.PVZResponse{
		ID:               pvz.ID.String(),
		RegistrationDate: pvz.RegistrationDate.String(),
		City:             string(pvz.Info.City),
	}
}

func ToPaginationFilterFromPaginationRequest(pag *dto.PaginationRequest) (*model.Filter, error) {
	var (
		filter model.Filter
		err    error
	)
	if len(pag.StartDate) != 0 {
		filter.StartDate, err = time.Parse(time.RFC3339, pag.StartDate)
		if err != nil {
			return nil, err
		}
	}

	if len(pag.EndDate) != 0 {
		filter.EndDate, err = time.Parse(time.RFC3339, pag.EndDate)
		if err != nil {
			return nil, err
		}
	}

	if len(pag.Limit) != 0 {
		filter.Limit, err = strconv.ParseUint(pag.Limit, 10, 64)
		if err != nil {
			return nil, err
		}

		if filter.Limit > 30 || filter.Limit < 1 {
			return nil, errors.New("limit must be between 1 and 30")
		}
	} else {
		filter.Limit = 10
	}

	var page uint64
	if len(pag.Page) != 0 {
		page, err = strconv.ParseUint(pag.Page, 10, 64)
		if err != nil {
			return nil, err
		}
		if page < 1 {
			return nil, errors.New("page must be greater or equal than 1")
		}
	} else {
		page = 1
	}

	filter.Offset = (page - 1) * filter.Limit
	return &filter, nil
}

func ToPVZListReceptionsResponse(pvzs []*model.PVZWithReceptions) []*dto.PVZListResponse {
	response := make([]*dto.PVZListResponse, 0, len(pvzs))
	for _, pvz := range pvzs {
		pvzResponse := dto.PVZResponse{
			ID:               pvz.PVZ.ID.String(),
			City:             string(pvz.PVZ.Info.City),
			RegistrationDate: pvz.PVZ.RegistrationDate.String(),
		}

		receptions := make([]dto.ReceptionsWithProducts, 0, len(pvz.Receptions))
		for _, rec := range pvz.Receptions {
			products := make([]dto.ProductResponse, 0, len(rec.Products))
			for _, pr := range rec.Products {
				products = append(products, *ToProductResponse(&pr))
			}
			receptions = append(receptions, dto.ReceptionsWithProducts{
				Reception: *ToReceptionResponse(&rec.Reception),
				Products:  products,
			})
		}

		response = append(response, &dto.PVZListResponse{
			PVZ:        pvzResponse,
			Receptions: receptions,
		})
	}
	return response
}

func ToProductResponse(pvz *model.Product) *dto.ProductResponse {
	return &dto.ProductResponse{
		ID:          pvz.ID.String(),
		DateTime:    pvz.CreatedAt.String(),
		Type:        string(pvz.Info.Type),
		ReceptionID: pvz.Info.ReceptionId.String(),
	}
}

func ToReceptionResponse(rec *model.Reception) *dto.ReceptionResponse {
	return &dto.ReceptionResponse{
		ID:       rec.ID.String(),
		PvzID:    rec.PvzId.String(),
		DateTime: rec.CreatedAt.String(),
		Status:   string(rec.Status),
	}
}

func ToProductPvzFromDTO(req *dto.ProductCreateRequest) (*model.ProductPVZ, error) {
	var (
		pvz model.ProductPVZ
		err error
	)
	pvz.PvzId, err = uuid.Parse(req.PvzID)
	if err != nil {
		return nil, err
	}

	pvz.Type = model.ProductType(req.Type)
	return &pvz, nil
}
