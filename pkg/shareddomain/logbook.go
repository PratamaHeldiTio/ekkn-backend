package shareddomain

import "backend-ekkn/modules/logbook/domain"

type LogbookRequest struct {
	PeriodID  string `json:"period_id"`
	StudentID string // get from jwt user
	Activity  string `json:"activity"`
	Image     string // get from filename form data
	date      string
}

type LogbookResponse struct {
	ID        string  `json:"id"`
	PeriodID  string  `json:"period_id"`
	StudentID string  `json:"student_id"`
	Activity  string  `json:"activity"`
	Image     string  `json:"image"`
	Radius    float64 `json:"radius"`
	Date      int64   `json:"date"`
	Submitted int64   `json:"submitted"`
}

func ToLogbookResponse(logbook domain.Logbook) LogbookResponse {
	return LogbookResponse{
		ID:        logbook.ID,
		PeriodID:  logbook.PeriodID,
		StudentID: logbook.StudentID,
		Activity:  logbook.Activity,
		Image:     logbook.Image,
		Radius:    logbook.Radius,
		Date:      logbook.Date,
		Submitted: logbook.Submitted,
	}
}
