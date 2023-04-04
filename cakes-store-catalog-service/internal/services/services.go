package services

import (
	"errors"

	"github.com/infinitss13/Cakes-store-catalog-service/entities"
)

type ServiceCatalog struct {
	Database DatabaseCatalog
}

func NewServiceCatalog(database DatabaseCatalog) *ServiceCatalog {
	return &ServiceCatalog{Database: database}
}

type DatabaseCatalog interface {
	GetCatalog() ([]entities.Cake, error)
}

func (s ServiceCatalog) GetCatalog() ([]entities.Cake, error) {
	cakes, err := s.Database.GetCatalog()
	if err != nil {
		return nil, err
	}
	if cakes == nil {
		return nil, errors.New("error : Catalog is empty, pinganite Stasa")
	}
	return cakes, nil
}
