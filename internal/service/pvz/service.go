package pvz

import (
	"github.com/biryanim/avito-tech-pvz/internal/repository"
	"github.com/biryanim/avito-tech-pvz/internal/service"
)

type serv struct {
	pvzRepository repository.PvzRepository
}

func NewService(pvzRepository repository.PvzRepository) service.PVZService {
	return &serv{
		pvzRepository: pvzRepository,
	}
}
