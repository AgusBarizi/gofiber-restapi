package mapper

import (
	"m3gaplazma/gofiber-restapi/model/domain"
	"m3gaplazma/gofiber-restapi/model/dto"
)

func ToUserResponse(user domain.User) dto.UserResponse {
	return dto.UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		IsActive: user.IsActive,
	}
}

func ToUserLoginResponse(user domain.User) dto.UserLoginResponse {
	return dto.UserLoginResponse{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		ApiToken: user.ApiToken,
	}
}

func ToUserResponses(users []domain.User) []dto.UserResponse {
	userResponses := []dto.UserResponse{}
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}
