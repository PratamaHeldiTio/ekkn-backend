package repository

import "backend-ekkn/modules/lecturer_registration/domain"

type LecturerRegistrationRepository interface {
	Create(registration domain.LecturerRegistration) error
	FindByPeriodLectureID(periodID, lectureID string) (domain.LecturerRegistration, error)
}
