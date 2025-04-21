package model

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrNoOpenReceptions       = errors.New("no open receptions")
	ErrNoProducts             = errors.New("no products")
	ErrNoSuchPvz              = errors.New("no such pvz")
	ErrLastReceptionNotClosed = errors.New("last reception not closed")
)

type PVZ struct {
	ID               uuid.UUID
	Info             PVZInfo
	RegistrationDate time.Time
}

type PVZInfo struct {
	City City
}

type Reception struct {
	ID        uuid.UUID
	CreatedAt time.Time
	PvzId     uuid.UUID
	Status    Status
}

type ProductPVZ struct {
	Type  ProductType
	PvzId uuid.UUID
}

type Product struct {
	ID        uuid.UUID
	Info      ProductInfo
	CreatedAt time.Time
}

type ProductInfo struct {
	Type        ProductType
	ReceptionId uuid.UUID
}

type PVZWithReceptions struct {
	PVZ        PVZ
	Receptions []ReceptionsWithProducts
}

type ReceptionsWithProducts struct {
	Reception Reception
	Products  []Product
}

type Filter struct {
	StartDate time.Time
	EndDate   time.Time
	Offset    uint64
	Limit     uint64
}

type City string

const (
	CityMoscow          City = "Москва"
	CitySaintPetersburg City = "Санкт-Петербург"
	CityKazan           City = "Казань"
)

func (c City) IsValid() bool {
	switch c {
	case CityMoscow, CitySaintPetersburg, CityKazan:
		return true
	}

	return false
}

type Status string

const (
	StatusInProgress Status = "in_progress"
	StatusClose      Status = "close"
)

func (s Status) IsValid() bool {
	switch s {
	case StatusInProgress, StatusClose:
		return true
	}
	return false
}

type ProductType string

const (
	ProductTypeElectronics ProductType = "электроника"
	ProductTypeClothing    ProductType = "одежда"
	ProductTypeShoes       ProductType = "обувь"
)

func (p ProductType) IsValid() bool {
	switch p {
	case ProductTypeElectronics, ProductTypeClothing, ProductTypeShoes:
		return true
	}

	return false
}
