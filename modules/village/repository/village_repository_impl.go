package repository

import (
	"backend-ekkn/modules/village/domain"
	"gorm.io/gorm"
)

type VillageRepositoryImpl struct {
	db *gorm.DB
}

func NeWVillageRepository(db *gorm.DB) VillageRespository {
	return &VillageRepositoryImpl{db}
}

func (repo *VillageRepositoryImpl) Create(village domain.Village) error {
	if err := repo.db.Create(&village).Error; err != nil {
		return err
	}

	return nil
}

func (repo *VillageRepositoryImpl) FindByPeriod(periodID, query string) ([]domain.Village, error) {
	var vilages []domain.Village
	if err := repo.db.Where("period_id = ?", periodID).
		Where("name ILIKE ? ", "%"+query+"%").
		Find(&vilages).Error; err != nil {
		return vilages, err
	}

	return vilages, nil
}

func (repo *VillageRepositoryImpl) Update(village domain.Village) error {
	if err := repo.db.Model(&village).Updates(village).Error; err != nil {
		return err
	}

	return nil
}

func (repo *VillageRepositoryImpl) FindById(ID string) (domain.Village, error) {
	var village domain.Village
	if err := repo.db.Where("id = ?", ID).Find(&village).Error; err != nil {
		return village, err
	}

	return village, nil
}

func (repo *VillageRepositoryImpl) Delete(village domain.Village) error {
	if err := repo.db.Delete(&village).Error; err != nil {
		return err
	}

	return nil
}
