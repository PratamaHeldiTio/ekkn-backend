package service

import (
	"backend-ekkn/modules/lecturer/domain"
	"backend-ekkn/pkg/shareddomain"
)

type LecturerService interface {
	CreateLecturer(request shareddomain.LecturerRequest) error
	UpdateLecturer(request shareddomain.LecturerRequest) error
	FindLecturerByID(ID string) (domain.Lecturer, error)
	FindAllLecturer(query string) ([]domain.Lecturer, error)
	LoginLecturer(request shareddomain.LecturerLogin) (domain.Lecturer, error)
	DeleteLecture(ID string) error
	ResetPassword(ID string) error
	ChangePassword(request shareddomain.ChangePasswordLecturerRequest) error
}
