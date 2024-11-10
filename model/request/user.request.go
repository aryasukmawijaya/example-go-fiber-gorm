package request

type UserCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type UserUpdateRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone"`
}

type UserEmailUpdateRequest struct {
	Email    string `json:"email" validate:"required"`
}