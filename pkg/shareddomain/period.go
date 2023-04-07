package shareddomain

type RequestPeriod struct {
	Semester              string `json:"semester" binding:"required,max=6"`
	TahunAjaran           string `json:"tahun_ajaran" binding:"required,max=10"`
	StatusRegisterStudent bool   `json:"status_register_student"`
	StatusRegisterLecture bool   `json:"status_register_lecture"`
	StatusRegisterGroup   bool   `json:"status_register_group" `
	Status                bool   `json:"status"`
	Start                 int64  `json:"start" binding:"required"`
	End                   int64  `json:"end" binding:"required"`
}
