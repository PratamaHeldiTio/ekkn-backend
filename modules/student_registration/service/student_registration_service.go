package service

import "backend-ekkn/pkg/shareddomain"

type StudentRegistrationService interface {
	CreateStudentRegistration(request shareddomain.RequestStudentRegistration) error
}
