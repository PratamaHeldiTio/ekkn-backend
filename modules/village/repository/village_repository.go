package repository

import "backend-ekkn/modules/village/domain"

type VillageRespository interface {
	Create(village domain.Village) error
	FindAll() ([]domain.Village, error)
	FindById(ID string) (domain.Village, error)
	Update(village domain.Village) error
}
