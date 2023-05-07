package service

import (
	"backend-ekkn/modules/village/domain"
	"backend-ekkn/pkg/shareddomain"
)

type VillageService interface {
	CreateVillage(request shareddomain.RequestVillage) error
	FindVillageByPeriod(periodID string) ([]domain.Village, error)
	UpdateVillage(request shareddomain.UpdateVillageRequest) error
	FindVillageById(ID string) (domain.Village, error)
	DeleteVillage(ID string) error
	AddDescVillage(request shareddomain.AddDescVillage) error
}
