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
		PeriodID:  request.PeriodID,
		Kecamatan: request.Kecamatan,
		Kabupaten: request.Kabupaten,
	}

	if err := service.repo.Create(village); err != nil {
		return err
	}

	return nil
}

func (service *VillageServiceImpl) FindVillageByPeriod(periodID string) ([]domain.Village, error) {
	villages, err := service.repo.FindByPeriod(periodID)
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

	village.Name = request.Name
	village.Kecamatan = request.Kecamatan
	village.Kabupaten = request.Kabupaten
	village.Latitude = request.Latitude
	village.Longitude = request.Longitude
	village.Status = request.Status

	if err := service.repo.Update(village); err != nil {
		return err
	}

	return nil
}

func (service *VillageServiceImpl) AddDescVillage(request shareddomain.AddDescVillage) error {
	village, err := service.repo.FindById(request.ID)
	if err != nil {
		return err
	}

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

func (service *VillageServiceImpl) DeleteVillage(ID string) error {
	village, err := service.FindVillageById(ID)
	if err != nil {
		return err
	}

	if err := service.repo.Delete(village); err != nil {
		return err
	}

	return nil
}
