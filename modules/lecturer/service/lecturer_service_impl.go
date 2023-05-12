package service

import (
	"backend-ekkn/modules/lecturer/domain"
	"backend-ekkn/modules/lecturer/repository"
	"backend-ekkn/pkg/shareddomain"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type LecturerServiceImpl struct {
	repo repository.LecturerRepository
}

func NewLecturerService(repo repository.LecturerRepository) LecturerService {
	return &LecturerServiceImpl{repo}
}

func (service *LecturerServiceImpl) CreateLecturer(request shareddomain.LecturerRequest) error {
	// generate password
	password, err := bcrypt.GenerateFromPassword([]byte(request.ID), bcrypt.MinCost)
	if err != nil {
		return err
	}

	lecturer := domain.Lecturer{
		ID:       request.ID,
		Name:     request.Name,
		Password: string(password),
	}

	if err := service.repo.Create(lecturer); err != nil {
		return err
	}

	return nil
}

func (service *LecturerServiceImpl) FindLecturerByID(ID string) (domain.Lecturer, error) {
	lecturer, err := service.repo.FindByID(ID)
	if err != nil {
		return lecturer, err
	}

	if lecturer.ID == "" {
		return lecturer, errors.New("data tidak ditemukan")
	}

	return lecturer, nil
}

func (service *LecturerServiceImpl) UpdateLecturer(request shareddomain.LecturerRequest) error {
	lecturer, err := service.FindLecturerByID(request.ID)
	if err != nil {
		return err
	}

	lecturer.Name = request.Name
	lecturer.Gender = request.Gender
	lecturer.Prodi = request.Prodi
	lecturer.Fakultas = request.Fakultas
	lecturer.MaduraLang = request.MaduraLang
	lecturer.Contact = request.Contact

	if err := service.repo.Update(lecturer); err != nil {
		return err
	}

	return nil
}

func (service *LecturerServiceImpl) DeleteLecture(ID string) error {
	lecturer, err := service.FindLecturerByID(ID)
	if err != nil {
		return err
	}

	if err := service.repo.Delete(lecturer); err != nil {
		return err
	}

	return nil
}

func (service *LecturerServiceImpl) ResetPassword(ID string) error {
	lecturer, err := service.FindLecturerByID(ID)
	if err != nil {
		return err
	}

	// hashing password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(lecturer.ID), bcrypt.MinCost)
	if err != nil {
		return err
	}

	// update password
	lecturer.Password = string(passwordHash)
	if err := service.repo.Update(lecturer); err != nil {
		return err
	}

	return nil
}

func (service *LecturerServiceImpl) FindAllLecturer() ([]domain.Lecturer, error) {
	lecturers, err := service.repo.FindAll()
	if err != nil {
		return lecturers, err
	}

	return lecturers, nil
}

func (service *LecturerServiceImpl) LoginLecturer(request shareddomain.LecturerLogin) (domain.Lecturer, error) {
	lecturer, err := service.FindLecturerByID(request.ID)
	if err != nil {
		return lecturer, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(lecturer.Password), []byte(request.Password)); err != nil {
		return lecturer, err
	}

	return lecturer, nil
}

func (service *LecturerServiceImpl) ChangePassword(request shareddomain.ChangePasswordLecturerRequest) error {
	lecturer, err := service.FindLecturerByID(request.ID)
	if err != nil {
		return err
	}

	// old password match with password db
	// check password is match
	if err := bcrypt.CompareHashAndPassword([]byte(lecturer.Password), []byte(request.OldPassword)); err != nil {
		return err
	}

	// check repeat password match with new password
	if request.NewPassword != request.RepeatNewPassword {
		return errors.New("password baru tidak match dengan ulangi password")
	}

	// hashing new password
	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.MinCost)
	if err != nil {
		return err
	}

	// update table student
	lecturer.Password = string(newPasswordHash)
	if err := service.repo.Update(lecturer); err != nil {
		return err
	}

	return nil
}
