package service

import (
	"backend-ekkn/modules/village/domain"
	"backend-ekkn/modules/village/repository"
	"backend-ekkn/pkg/shareddomain"
	"errors"
)

type VillageServiceImpl struct {
	repo repository.VillageRespository
}

func NewVillageService(repo repository.VillageRespository) VillageService {
	return &VillageServiceImpl{repo}
}

func (service *VillageServiceImpl) CreateVillage(request shareddomain.RequestVillage) error {
	// type requesst to domain
	village := domain.Village{
		Name:      request.Name,
		Kecamatan: request.Kecamatan,
		Kabupaten: request.Kabupaten,
	}

	if err := service.repo.Create(village); err != nil {
		return err
	}

	return nil
}

func (service *VillageServiceImpl) FindAllVillage() ([]domain.Village, error) {
	villages, err := service.repo.FindAll()
	if err != nil {
		return villages, err
	}

	return villages, err
}

func (service *VillageServiceImpl) UpdateVillage(request shareddomain.UpdateVillageRequest) error {
	village, err := service.repo.FindById(request.ID)
	if err != nil {
		return err
	}

	village.Status = request.Status
	village.Strength = request.Strength
	village.Weakness = request.Weakness
	village.Oportunities = request.Oportunities
	village.Threats = request.Threats

	if err := service.repo.Update(village); err != nil {
		return err
	}

	return nil
}

func (service *VillageServiceImpl) FindVillageById(ID string) (domain.Village, error) {
	// get group by student id and period id
	village, err := service.repo.FindById(ID)
	if err != nil {
		return village, nil
	}

	// cek isExist
	if village.ID == "" {
		return village, errors.New("desa tidak dapat ditemukan")
	}

	return village, nil
}
