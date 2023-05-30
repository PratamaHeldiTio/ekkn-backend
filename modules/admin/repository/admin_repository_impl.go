package repository

import (
	"backend-ekkn/modules/admin/domain"
	"gorm.io/gorm"
)

type AdminRepositoryImpl struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &AdminRepositoryImpl{db}
}

func (repo *AdminRepositoryImpl) Create(admin domain.Admin) error {
	if err := repo.db.Create(&admin).Error; err != nil {
		return err
	}

	return nil
}

func (repo *AdminRepositoryImpl) FindByUsername(username string) (domain.Admin, error) {
	var admin domain.Admin
	if err := repo.db.Where("username = ?", username).Find(&admin).Error; err != nil {
		return admin, err
	}

	return admin, nil
}

func (repo *AdminRepositoryImpl) Delete(admin domain.Admin) error {
	if err := repo.db.Delete(&admin).Error; err != nil {
		return err
	}

	return nil
}
