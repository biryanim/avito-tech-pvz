package model

import (
	"github.com/google/uuid"
	"time"
)

type Pvz struct {
	ID               uuid.UUID
	RegistrationDate time.Time
	City             City
}

type Reception struct {
	ID       uuid.UUID
	DateTime time.Time
	PvzId    uuid.UUID
	Status   Status
}

type Product struct {
	ID          uuid.UUID
	Date        time.Time
	Type        ProductType
	ReceptionId uuid.UUID
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
