package service

import (
	"backend-ekkn/modules/lecturer_registration/domain"
	"backend-ekkn/pkg/shareddomain"
)

type LecturerRegistrationService interface {
	LecturerRegistration(request shareddomain.LecturerRegistrationRequest) error
	FindLecturerRegistrationByLectureID(lectureID string) ([]domain.LecturerRegistration, error)
	ValidationLecturerRegistration(request shareddomain.ValidationLectureRegistrationRequest) error
	FindLecturerRegistrationByID(ID string) (domain.LecturerRegistration, error)
	FindLecturerRegistrationByPeriod(ID string) ([]domain.LecturerRegistration, error)
}
