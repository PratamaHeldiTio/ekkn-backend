package shareddomain

import "github.com/google/uuid"

type RequestStudentRegistration struct {
	ID       uuid.UUID `json:"student_registration_id"`
	PeriodID uuid.UUID `json:"period_id" binding:"required"`
	Nim      string    `binding:"max=14"`
	Status   string    `json:"status" binding:"max=14"`
}
