package service

import (
	"backend-ekkn/modules/village/domain"
	"backend-ekkn/pkg/shareddomain"
)

type VillageService interface {
	CreateVillage(request shareddomain.RequestVillage) error
	FindAllVillage() ([]domain.Village, error)
	UpdateVillage(request shareddomain.UpdateVillageRequest) error
	FindVillageById(ID string) (domain.Village, error)
}
