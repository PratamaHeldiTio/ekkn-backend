package service

import (
	"backend-ekkn/modules/lecturer_registration/domain"
	"backend-ekkn/pkg/shareddomain"
)

type LecturerRegistrationService interface {
	LecturerRegistration(request shareddomain.LecturerRegistrationRequest) error
	FindLecturerRegistrationByLectureID(lectureID string) ([]domain.LecturerRegistration, error)
}
