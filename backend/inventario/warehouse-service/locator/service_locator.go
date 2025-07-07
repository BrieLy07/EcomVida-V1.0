package locator

import (
	"warehouse-service/repository"
)

type ServiceLocator struct {
	warehouseRepo repository.WarehouseRepository
}

func (s *ServiceLocator) GetWarehouseRepository() repository.WarehouseRepository {
	if s.warehouseRepo == nil {
		s.warehouseRepo = repository.NewWarehouseRepository()
	}
	return s.warehouseRepo
}

var Locator = &ServiceLocator{}
