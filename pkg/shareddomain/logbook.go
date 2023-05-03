package shareddomain

import (
	"backend-ekkn/modules/logbook/domain"
	"mime/multipart"
)

type LogbookRequest struct {
	FileImage *multipart.FileHeader `form:"image" binding:"required"`
	PeriodID  string                `form:"period_id" binding:"required"`
	StudentID string                // get from jwt user
	GroupID   string                `form:"group_id" binding:"required"`
	Activity  string                `form:"activity" binding:"required"`
	Image     string                // get from filename form data
	Date      string                `form:"date" binding:"required"`
	Latitude  float64               `form:"latitude" binding:"latitude"`
	Longitude float64               `form:"longitude" binding:"longitude"`
}

type LogbookResponse struct {
	ID        string `json:"id"`
	PeriodID  string `json:"period_id"`
	StudentID string `json:"student_id"`
	Activity  string `json:"activity"`
	Image     string `json:"image"`
	Radius    int    `json:"radius"`
	Date      int64  `json:"date"`
	Submitted int64  `json:"submitted"`
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
