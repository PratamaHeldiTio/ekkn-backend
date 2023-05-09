package shareddomain

import (
	"backend-ekkn/modules/student_registration/domain"
	"github.com/google/uuid"
)

type RequestStudentRegistration struct {
	ID       uuid.UUID `json:"student_registration_id"`
	PeriodID string    `json:"period_id" binding:"required"`
	Nim      string    `binding:"max=14"`
	Status   string    `json:"status" binding:"max=14"`
}

type StudentRegistrationURI struct {
	PeriodID  string `uri:"periodID" binding:"required"`
	StudentID string `uri:"studentID" binding:"required"`
}

type ResponseStudentRegistrationByNim struct {
	ID          string `json:"student_registration_id"`
	PeriodID    string `json:"period_id"`
	Semester    string `json:"semester"`
	TahunAjaran string `json:"tahun_ajaran"`
	StudentID   string `json:"nim"`
	Name        string `json:"name"`
	Prodi       string `json:"prodi"`
	Fakultas    string `json:"fakultas"`
	Status      string `json:"status"`
	CreatedAt   int64  `json:"created_at"`
}

type ResponseStudentRegistrationByNimPeriodID struct {
	Status string `json:"status"`
	Group  string `json:"group"`
}

type ResponseRegisteredStudents struct {
	PeriodID    string `json:"period_id"`
	Semester    string `json:"semester"`
	TahunAjaran string `json:"tahun_ajaran"`
}

type UpdateStudentRegistrationRequest struct {
	ID     string
	Status string `json:"status" binding:"required"`
}

type StudentRegistrationPeriodResponse struct {
	ID        string `json:"id"`
	StudentID string `json:"nim"`
	Name      string `json:"name"`
	Prodi     string `json:"prodi"`
	Status    string `json:"status"`
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

func ToResponseStudentRegistrationByNimPeriodID(registration domain.StudentRegistration) ResponseStudentRegistrationByNimPeriodID {
	registeredStudents := ResponseStudentRegistrationByNimPeriodID{
		Status: registration.Status,
		Group:  registration.Group,
	}
	return registeredStudents
}

func ToStudentRegistrationPeriod(registration domain.StudentRegistration) StudentRegistrationPeriodResponse {
	registeredStudents := StudentRegistrationPeriodResponse{
		ID:        registration.ID,
		StudentID: registration.StudentID,
		Name:      registration.Student.Name,
		Prodi:     registration.Student.Prodi,
		Status:    registration.Status,
	}
	return registeredStudents
}
