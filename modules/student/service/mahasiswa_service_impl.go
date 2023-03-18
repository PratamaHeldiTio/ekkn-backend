package service

import (
	"backend-ekkn/modules/student/domain"
	"backend-ekkn/modules/student/repository"
	"backend-ekkn/pkg/shareddomain"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type StudentServiceImpl struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &StudentServiceImpl{repo}
}

func (service *StudentServiceImpl) CreateStudent(request shareddomain.CreateStudentRequest) (domain.Student, error) {

	createdAt := time.Now().Unix()
	updateAt := createdAt
	Student := domain.Student{
		Nim:       request.Nim,
		Name:      request.Name,
		Prodi:     request.Prodi,
		Fakultas:  request.Fakultas,
		CreatedAt: createdAt,
		UpdateAt:  updateAt,
		Role:      "none",
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Nim), bcrypt.MinCost)
	if err != nil {
		return Student, err
	}

	Student.Password = string(passwordHash)
	Student, err = service.repo.Save(Student)
	if err != nil {
		return Student, err
	}

	return Student, nil
}
