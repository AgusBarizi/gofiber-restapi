package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"m3gaplazma/gofiber-restapi/exception"
	"m3gaplazma/gofiber-restapi/model/dto"
)

func CreateUserValidation(request dto.CreateUserRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Errors: err,
		})
	}
}

func UserLoginValidation(request dto.UserLoginRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Errors: err,
		})
	}
}

func ActivateUserValidation(request dto.ActivateUserRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.IsActive, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Errors: err,
		})
	}
}

func ChangeUserPasswordValidation(request dto.ChangeUserPasswordRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.OldPassword, validation.Required),
		validation.Field(&request.NewPassword, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Errors: err,
		})
	}
}
