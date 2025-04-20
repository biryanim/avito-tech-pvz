package pvz

import "github.com/biryanim/avito-tech-pvz/internal/service"

type Implementation struct {
	pvzService service.PVZService
}

func NewImplementation(pvzService service.PVZService) *Implementation {
	return &Implementation{
		pvzService: pvzService,
	}
}
