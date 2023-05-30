package shareddomain

type AdminRequest struct {
	Username string `json:"username" binding:"required,max=50"`
	Password string `json:"password" binding:"required"`
}

type DeleteAdminRequest struct {
	Username string `json:"username" binding:"required,max=50"`
}
