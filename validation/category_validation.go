package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"m3gaplazma/gofiber-restapi/exception"
	"m3gaplazma/gofiber-restapi/model/dto"
)

func CreateCategoryValidation(request dto.CreateCategoryRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.Name, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Errors: err,
		})
	}
}

func UpdateCategoryValidation(request dto.UpdateCategoryRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Name, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Errors: err,
		})
	}
}
