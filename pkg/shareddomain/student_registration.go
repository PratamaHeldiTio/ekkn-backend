package shareddomain

import (
	"backend-ekkn/modules/student_registration/domain"
	"github.com/google/uuid"
)

type RequestStudentRegistration struct {
	ID       uuid.UUID `json:"student_registration_id"`
	PeriodID uuid.UUID `json:"period_id" binding:"required"`
	Nim      string    `binding:"max=14"`
	Status   string    `json:"status" binding:"max=14"`
}

type ResponseStudentRegistrationByNim struct {
	ID          uuid.UUID `json:"student_registration_id"`
	PeriodID    uuid.UUID `json:"period_id"`
	Semester    string    `json:"semester"`
	TahunAjaran string    `json:"tahun_ajaran"`
	StudentID   string    `json:"nim"`
	Name        string    `json:"name"`
	Prodi       string    `json:"prodi"`
	Fakultas    string    `json:"fakultas"`
	Status      string    `json:"status"`
	CreatedAt   int64     `json:"created_at"`
}

type ResponseRegisteredStudents struct {
	PeriodID    uuid.UUID `json:"period_id"`
	Semester    string    `json:"semester"`
	TahunAjaran string    `json:"tahun_ajaran"`
}

func ToResponRegiteredStudent(registration domain.StudentRegistration) ResponseStudentRegistrationByNim {
	studentRegistered := ResponseStudentRegistrationByNim{
		ID:          registration.ID,
		PeriodID:    registration.PeriodID,
		Semester:    registration.Period.Semester,
		TahunAjaran: registration.Period.TahunAjaran,
		StudentID:   registration.StudentID,
		Name:        registration.Student.Name,
		Prodi:       registration.Student.Prodi,
		Fakultas:    registration.Student.Fakultas,
		Status:      registration.Status,
		CreatedAt:   registration.Student.CreatedAt,
	}

	return studentRegistered
}

func ToResponseRegisteredStudents(registration domain.StudentRegistration) ResponseRegisteredStudents {
	registeredStudents := ResponseRegisteredStudents{
		PeriodID:    registration.PeriodID,
		Semester:    registration.Period.Semester,
		TahunAjaran: registration.Period.TahunAjaran,
	}

	return registeredStudents
}
