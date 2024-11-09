package request

type UserCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}
