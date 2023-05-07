package repository

import "backend-ekkn/modules/village/domain"

type VillageRespository interface {
	Create(village domain.Village) error
	FindByPeriod(periodID string) ([]domain.Village, error)
	FindById(ID string) (domain.Village, error)
	Update(village domain.Village) error
	Delete(village domain.Village) error
}
