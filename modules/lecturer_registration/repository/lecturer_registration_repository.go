package repository

import (
	"backend-ekkn/modules/lecturer_registration/domain"
)

type LecturerRegistrationRepository interface {
	Create(registration domain.LecturerRegistration) error
	FindByPeriodLectureID(periodID, lectureID string) (domain.LecturerRegistration, error)
	FindByLectureID(lecturerID string) ([]domain.LecturerRegistration, error)
	FindByID(ID string) (domain.LecturerRegistration, error)
	Update(registration domain.LecturerRegistration) error
	FindByPeriod(ID, query string) ([]domain.LecturerRegistration, error)
}
