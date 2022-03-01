package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"m3gaplazma/gofiber-restapi/exception"
	"m3gaplazma/gofiber-restapi/model/dto"
)

func CreateProductValidation(request dto.CreateProductRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Stock, validation.Required, validation.Min(0)),
		validation.Field(&request.Sku, validation.Required),
		validation.Field(&request.Price, validation.Required, validation.Min(100)),
	)
	if err != nil {
		panic(exception.ValidationError{
			Errors: err,
		})
	}
}

func UpdateProductValidation(request dto.UpdateProductRequest) {
	err := validation.ValidateStruct(
		&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Stock, validation.Required, validation.Min(0)),
		validation.Field(&request.Sku, validation.Required),
		validation.Field(&request.Price, validation.Required, validation.Min(100)),
	)
	if err != nil {
		panic(exception.ValidationError{
			Errors: err,
		})
	}
}
