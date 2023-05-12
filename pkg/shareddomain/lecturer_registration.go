package shareddomain

type LecturerRegistrationRequest struct {
	PeriodID   string `json:"period_id" binding:"required"`
	LecturerID string
}
