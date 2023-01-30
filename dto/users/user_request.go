package usersdto


type UpdateUserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	// Roles    string    `json:"roles" form:"roles" validate:"required"`
}
