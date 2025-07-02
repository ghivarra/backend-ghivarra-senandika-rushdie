package auth

type UserRegister struct {
	Username             string `json:"username" binding:"required,min=6,max=50"`
	Password             string `json:"password" binding:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
	Email                string `json:"email" binding:"required,email"`
	Name                 string `json:"name" binding:"required,max=100"`
	UserRoleID           int    `json:"user_role_id" binding:"required,numeric"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
