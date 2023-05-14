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

type LecturerRegistrationApproveResponse struct {
	PeriodID    string `json:"period_id"`
	Semester    string `json:"semester"`
	TahunAjaran string `json:"tahun_ajaran"`
}

type ValidationLectureRegistrationRequest struct {
	ID     string
	Status string `json:"status" binding:"required"`
}

type LecturerRegistrationByPeriodResponse struct {
	ID         string `json:"id"`
	LecturerID string `json:"lecturer_id"`
	Name       string `json:"name"`
	Prodi      string `json:"prodi"`
	Fakultas   string `json:"fakultas"`
	Status     string `json:"status"`
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

func ToLecturerRegistrationByPeriod(lecturerRegistration domain.LecturerRegistration) LecturerRegistrationByPeriodResponse {
	return LecturerRegistrationByPeriodResponse{
		ID:         lecturerRegistration.ID,
		LecturerID: lecturerRegistration.LecturerID,
		Name:       lecturerRegistration.Lecturer.Name,
		Prodi:      lecturerRegistration.Lecturer.Prodi,
		Fakultas:   lecturerRegistration.Lecturer.Fakultas,
		Status:     lecturerRegistration.Status,
	}
}

func ToLecturerRegistrationApprove(lecturerRegistration domain.LecturerRegistration) LecturerRegistrationApproveResponse {
	return LecturerRegistrationApproveResponse{
		PeriodID:    lecturerRegistration.PeriodID,
		Semester:    lecturerRegistration.Period.Semester,
		TahunAjaran: lecturerRegistration.Period.TahunAjaran,
	}
}
