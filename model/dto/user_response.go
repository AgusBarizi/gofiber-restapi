package dto

type UserResponse struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive int    `json:"is_active"`
}

type UserLoginResponse struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	ApiToken string `json:"api_token"`
}
