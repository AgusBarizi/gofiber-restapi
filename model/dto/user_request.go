package dto

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ActivateUserRequest struct {
	Id       int64 `json:"id"`
	IsActive int   `json:"is_active"`
}

type ChangeUserPasswordRequest struct {
	Id          int64  `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
