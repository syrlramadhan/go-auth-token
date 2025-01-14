package dto

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Pass  string `json:"password"`
}

type LoginUserRequest struct {
	Email string `json:"email"`
	Pass  string `json:"password"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
