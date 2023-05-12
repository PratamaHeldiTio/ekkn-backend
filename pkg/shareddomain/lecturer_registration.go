package shareddomain

import "backend-ekkn/modules/lecturer_registration/domain"

type LecturerRegistrationRequest struct {
	PeriodID   string `json:"period_id" binding:"required"`
	LecturerID string
}

type LecturerRegistrationHistoryResponse struct {
	ID          string `json:"id"`
	Semester    string `json:"semester"`
	TahunAjaran string `json:"tahun_ajaran"`
	Status      string `json:"status"`
	CreatedAt   int64  `json:"created_at"`
}

func ToLecturerRegistrationHistory(lecturerRegistration domain.LecturerRegistration) LecturerRegistrationHistoryResponse {
	return LecturerRegistrationHistoryResponse{
		ID:          lecturerRegistration.ID,
		Semester:    lecturerRegistration.Period.Semester,
		TahunAjaran: lecturerRegistration.Period.TahunAjaran,
		Status:      lecturerRegistration.Status,
		CreatedAt:   lecturerRegistration.CreatedAt,
	}
}
