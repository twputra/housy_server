package usersdto

type UserResponse struct {
  ID       int    `json:"id"`
  Name     string `json:"name" form:"name" validate:"required"`
  Email     string `json:"email" form:"email" validate:"required"`
  Username    string `json:"username" form:"username" validate:"required"`
  Password string `json:"password" form:"password" validate:"required"`
  Gender string `json:"gender" form:"gender" validate:"required"`
  ListAsID int `json:"listasid" form:"listasid" validate:"required"`
  Address string `json:"address" form:"address" validate:"required"`
}