package service

import (
	"backend-ekkn/modules/admin/domain"
	"backend-ekkn/modules/admin/repository"
	"backend-ekkn/pkg/shareddomain"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type AdminServiceImpl struct {
	repo repository.AdminRepository
}

func NewAdminRepository(repo repository.AdminRepository) AdminService {
	return &AdminServiceImpl{repo}
}

func (service *AdminServiceImpl) CreateAdmin(request shareddomain.AdminRequest) error {
	// hashing password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	admin := domain.Admin{
		Username: request.Username,
		Password: string(passwordHash),
	}

	if err := service.repo.Create(admin); err != nil {
		return err
	}

	return nil
}

func (service *AdminServiceImpl) LoginAdmin(request shareddomain.AdminRequest) (domain.Admin, error) {
	admin, err := service.repo.FindByUsername(request.Username)
	if err != nil {
		return admin, err
	}

	if admin.Username == "" {
		return admin, errors.New("Data tidak ditemukan")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(request.Password)); err != nil {
		return admin, err
	}

	return admin, err
}

func (service *AdminServiceImpl) DeleteAdmin(username string) error {
	admin, err := service.repo.FindByUsername(username)
	if err != nil {
		return err
	}

	if admin.Username == "" {
		return errors.New("Data tidak ditemukan")
	}

	if err := service.repo.Delete(admin); err != nil {
		return err
	}

	return nil
}
