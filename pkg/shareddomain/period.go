package shareddomain

import (
	"backend-ekkn/modules/period/domain"
	"github.com/google/uuid"
)

type RequestPeriod struct {
	ID                    uuid.UUID `json:"id_period"`
	Semester              string    `json:"semester" binding:"required,max=6"`
	TahunAjaran           string    `json:"tahun_ajaran" binding:"required,max=10"`
	StatusRegisterStudent string    `json:"status_register_student" binding:"required,max=5"`
	StatusRegisterLecture string    `json:"status_register_lecture" binding:"required,max=5"`
	StatusRegisterGroup   string    `json:"status_register_group" binding:"required,max=5"`
	Status                string    `json:"status" binding:"required,max=5"`
	Start                 int64     `json:"start" binding:"required"`
	End                   int64     `json:"end" binding:"required"`
}

type ResponsePeriod struct {
	ID                    uuid.UUID `json:"period_id"`
	Semester              string    `json:"semester"`
	TahunAjaran           string    `json:"tahun_ajaran"`
	StatusRegisterStudent string    `json:"status_register_student"`
	StatusRegisterLecture string    `json:"status_register_lecture"`
	StatusRegisterGroup   string    `json:"status_register_group" `
	Status                string    `json:"status"`
	Start                 int64     `json:"start"`
	End                   int64     `json:"end"`
	CreatedAt             int64     `json:"created_at"`
	UpdatedAt             int64     `json:"updated_at"`
}

type ResponsePeriodStudent struct {
	ID          uuid.UUID `json:"period_id"`
	Semester    string    `json:"semester"`
	TahunAjaran string    `json:"tahun_ajaran"`
}

func ToResponsePeriod(period domain.Period) ResponsePeriod {
	periodResponse := ResponsePeriod{
		ID:                    period.ID,
		Semester:              period.Semester,
		TahunAjaran:           period.TahunAjaran,
		StatusRegisterStudent: period.StatusRegisterStudent,
		StatusRegisterLecture: period.StatusRegisterLecture,
		StatusRegisterGroup:   period.StatusRegisterGroup,
		Status:                period.Status,
		Start:                 period.Start,
		End:                   period.End,
		CreatedAt:             period.CreatedAt,
		UpdatedAt:             period.UpdatedAt,
	}

	return periodResponse
}

func ToResponsePeriodStudent(period domain.Period) ResponsePeriodStudent {
	periodResponse := ResponsePeriodStudent{
		ID:          period.ID,
		Semester:    period.Semester,
		TahunAjaran: period.TahunAjaran,
	}

	return periodResponse
}
