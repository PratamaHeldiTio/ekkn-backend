package service

import (
	"backend-ekkn/modules/student/domain"
	"backend-ekkn/modules/student/repository"
	"backend-ekkn/pkg/shareddomain"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type StudentServiceImpl struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &StudentServiceImpl{repo}
}

func (service *StudentServiceImpl) CreateStudent(request shareddomain.CreateStudent) error {

	student := domain.Student{
		Nim:      request.Nim,
		Name:     request.Name,
		Prodi:    request.Prodi,
		Fakultas: request.Fakultas,
	}

	// hashing password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Nim), bcrypt.MinCost)
	if err != nil {
		return err
	}

	student.Password = string(passwordHash)
	err = service.repo.Create(student)
	if err != nil {
		return err
	}
	return nil
}

func (service *StudentServiceImpl) FindStudentByNim(nim string) (domain.Student, error) {
	student, err := service.repo.FindByNim(nim)
	if err != nil {
		return student, err
	}

	if student.Nim == "" {
		return student, errors.New("No student found on that nim")
	}

	return student, nil
}

func (service *StudentServiceImpl) LoginStudent(request shareddomain.LoginStudent) (domain.Student, error) {
	// check nim request is exist
	student, err := service.repo.FindByNim(request.Nim)

	if err != nil {
		return student, err
	}

	if student.Nim == "" {
		return student, errors.New("No student found on that nim")
	}

	// check password is match
	err = bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(request.Password))
	if err != nil {
		return student, err
	}
	return student, nil
}

// get all student use repository
func (service *StudentServiceImpl) FindAllStudent(query string) ([]domain.Student, error) {
	students, err := service.repo.FindAll(query)
	if err != nil {
		return students, err
	}

	return students, nil
}

func (service *StudentServiceImpl) UpdateStudent(request shareddomain.UpdateStudent) error {
	// cek nim isExist
	student, err := service.repo.FindByNim(request.Nim)
	if err != nil {
		return err
	}

	if student.Nim == "" {
		return errors.New("No student found on that nim")
	}

	student = domain.Student{
		Nim:        request.Nim,
		Name:       request.Name,
		Prodi:      request.Prodi,
		Fakultas:   request.Fakultas,
		Gender:     request.Gender,
		MaduraLang: request.MaduraLang,
	}
	err = service.repo.Update(student)
	if err != nil {
		return err
	}

	return nil
}

func (service *StudentServiceImpl) DeleteStudent(nim string) error {
	student, err := service.repo.FindByNim(nim)
	if err != nil {
		return err
	}

	if student.Nim == "" {
		return errors.New("No student found on that nim")
	}

	if err := service.repo.Delete(student); err != nil {
		return err
	}

	return nil
}

func (service *StudentServiceImpl) ChangePassword(request shareddomain.ChangePasswordRequest) error {
	student, err := service.FindStudentByNim(request.Nim)
	if err != nil {
		return err
	}

	// old password match with password db
	// check password is match
	if err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(request.OldPassword)); err != nil {
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
	student.Password = string(newPasswordHash)
	if err := service.repo.Update(student); err != nil {
		return err
	}

	return nil
}

func (service *StudentServiceImpl) ResetPassword(studentID string) error {
	student, err := service.FindStudentByNim(studentID)
	if err != nil {
		return err
	}

	// hashing password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(studentID), bcrypt.MinCost)
	if err != nil {
		return err
	}

	// update password
	student.Password = string(passwordHash)
	if err := service.repo.Update(student); err != nil {
		return err
	}

	return nil
}
