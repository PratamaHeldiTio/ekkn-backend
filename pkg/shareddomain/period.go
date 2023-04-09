package shareddomain

import (
	"backend-ekkn/modules/period/domain"
	"github.com/google/uuid"
)

type RequestPeriod struct {
	ID                    uuid.UUID `json:"id_period"`
	Semester              string    `json:"semester" binding:"required,max=6"`
	TahunAjaran           string    `json:"tahun_ajaran" binding:"required,max=10"`
	StatusRegisterStudent bool      `json:"status_register_student"`
	StatusRegisterLecture bool      `json:"status_register_lecture"`
	StatusRegisterGroup   bool      `json:"status_register_group" `
	Status                bool      `json:"status"`
	Start                 int64     `json:"start" binding:"required"`
	End                   int64     `json:"end" binding:"required"`
}

type ResponsePeriod struct {
	ID                    uuid.UUID `json:"id_period"`
	Semester              string    `json:"semester"`
	TahunAjaran           string    `json:"tahun_ajaran"`
	StatusRegisterStudent bool      `json:"status_register_student"`
	StatusRegisterLecture bool      `json:"status_register_lecture"`
	StatusRegisterGroup   bool      `json:"status_register_group" `
	Status                bool      `json:"status"`
	Start                 int64     `json:"start"`
	End                   int64     `json:"end"`
	CreatedAt             int64     `json:"created_at"`
	UpdatedAt             int64     `json:"updated_at"`
}

type ResponsePeriodBasic struct {
	ID          uuid.UUID `json:"id_period"`
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

func ToResponsePeriodBasic(period domain.Period) ResponsePeriodBasic {
	periodResponse := ResponsePeriodBasic{
		ID:          period.ID,
		Semester:    period.Semester,
		TahunAjaran: period.TahunAjaran,
	}

	return periodResponse
}
