package authdto

type RegisterResponse struct {
	Username string `gorm:"type: varchar(255)" json:"username"`
	Message  string `gorm:"type: varchar(255)" json:"message"`
}

type LoginResponse struct {
	Username string `gorm:"type: varchar(255)" json:"username"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}