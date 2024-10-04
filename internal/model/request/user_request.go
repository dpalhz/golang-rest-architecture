package request

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRegister struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserUpdate struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty" validate:"omitempty,email"`
	Photo *string `json:"photo,omitempty"`
}
