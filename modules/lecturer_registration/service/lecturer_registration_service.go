package service

import "backend-ekkn/pkg/shareddomain"

type LecturerRegistrationService interface {
	LecturerRegistration(request shareddomain.LecturerRegistrationRequest) error
}
