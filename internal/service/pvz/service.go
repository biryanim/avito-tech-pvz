package pvz

import (
	"github.com/biryanim/avito-tech-pvz/internal/client/db"
	"github.com/biryanim/avito-tech-pvz/internal/repository"
	"github.com/biryanim/avito-tech-pvz/internal/service"
)

type serv struct {
	pvzRepository repository.PvzRepository
	txManager     db.TxManager
}

func NewService(pvzRepository repository.PvzRepository, txManager db.TxManager) service.PVZService {
	return &serv{
		pvzRepository: pvzRepository,
		txManager:     txManager,
	}
}
