package user

// "gorm.io/gorm"

type UserCreateResponse struct {
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Saldo       string `json:"saldo"`

	// Roles    bool   `json:"roles"`
	// Image    string `json:"image"`
}
type UserUpdateResponse struct {
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Saldo       string `json:"saldo"`
	// Roles    bool   `json:"roles"`
	// Image    string `json:"image"`
}
type UserGetByIdResponse struct {
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Saldo       string `json:"saldo"`

	// Roles    bool   `json:"roles"`
	// Image    string `json:"image"`
}

//=========================================================

// =================== Create User Request =======================
type CreateUserRequestFormat struct {
	Name        string `json:"name" form:"name" validate:"required,min=3,max=25,excludesall=!@#?^$*()_+-=0123456789%&"`
	Email       string `json:"email" form:"email" validate:"required,email"`
	Password    string `json:"password" form:"password" validate:"required,min=3,max=20"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	// Image    string `json:"image" form:"image"`
}

// =================== Update User Request =======================
type UpdateUserRequestFormat struct {
	Name        string `json:"name" form:"name" validate:"omitempty,min=3,max=25,excludesall=!@#?^$*()_+-=0123456789%&"`
	Email       string `json:"email" form:"email" validate:"omitempty,email"`
	Password    string `json:"password" form:"password" validate:"omitempty,min=3,max=20"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"omitempty"`
	// Image       string `json:"image" form:"image"`
}
