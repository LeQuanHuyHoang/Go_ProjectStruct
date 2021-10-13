package model

type (
	User struct {
		BaseModel
		Email    string `json:"email" gorm:"column:email"`
		Password string `json:"password"`
	}

	UserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
