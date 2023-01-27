package usersdto

type CreateUserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Username    string `json:"username" form:"username" validate:"required"`
	Email string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	ListAsID int `json:"listasid" form:"listasid" validate:"required"`
	Gender string `json:"gender" form:"gender" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
  }
  
  type UpdateUserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Username    string `json:"username" form:"username" validate:"required"`
	Email string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	ListAsID int `json:"listasid" form:"listasid" validate:"required"`
	Gender string `json:"gender" form:"gender" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
  }