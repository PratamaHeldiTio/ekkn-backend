package shareddomain

import (
	"backend-ekkn/modules/period/domain"
)

type RequestPeriod struct {
	ID                        string
	Semester                  string `json:"semester" binding:"required,max=6"`
	TahunAjaran               string `json:"tahun_ajaran" binding:"required,max=10"`
	StudentRegistrationStatus string `json:"student_registration_status" binding:"max=5"`
	LectureRegistrationStatus string `json:"lecture_registration_status" binding:"max=5"`
	GroupRegistrationStatus   string `json:"group_registration_status" binding:"max=5"`
	Status                    string `json:"status" binding:"max=5"`
	Start                     string `json:"start" binding:"required"`
	End                       string `json:"end" binding:"required"`
}

type ResponsePeriod struct {
	ID                    string `json:"id"`
	Semester              string `json:"semester"`
	TahunAjaran           string `json:"tahun_ajaran"`
	StatusRegisterStudent string `json:"student_registration_status"`
	StatusRegisterLecture string `json:"lecture_registration_status"`
	StatusRegisterGroup   string `json:"group_registration_status" `
	Status                string `json:"status"`
	Start                 int64  `json:"start"`
	End                   int64  `json:"end"`
	CreatedAt             int64  `json:"created_at"`
	UpdatedAt             int64  `json:"updated_at"`
}

type ResponsePeriodBasic struct {
	ID          string `json:"period_id"`
	Semester    string `json:"semester"`
	TahunAjaran string `json:"tahun_ajaran"`
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

func ToResponsePeriodStudent(period domain.Period) ResponsePeriodBasic {
	periodResponse := ResponsePeriodBasic{
		ID:          period.ID,
		Semester:    period.Semester,
		TahunAjaran: period.TahunAjaran,
	}

	return periodResponse
}
