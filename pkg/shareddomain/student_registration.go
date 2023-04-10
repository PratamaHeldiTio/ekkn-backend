package shareddomain

import "github.com/google/uuid"

type RequestStudentRegistration struct {
	ID       uuid.UUID `json:"id_student_registration"`
	PeriodID uuid.UUID `json:"id_period" binding:"required"`
	Nim      string    `binding:"max=14"`
	Status   string    `json:"status" binding:"max=14"`
}
