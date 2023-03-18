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

func (repo *AdminRepositoryImpl) Save(admin domain.Admin) (domain.Admin, error) {
	if err := repo.db.Create(&admin).Error; err != nil {
		return admin, err
	} else {
		return admin, nil
	}
}
