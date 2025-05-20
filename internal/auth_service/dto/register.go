package dto

type RegisterRequest struct {
    FullName string `json:"full_name" validate:"required"`
    Email    string `json:"email"      validate:"required,email"`
    Password string `json:"password"   validate:"required,min=6"`
}


type RegisterResponse struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}