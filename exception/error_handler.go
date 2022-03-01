package exception

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"m3gaplazma/gofiber-restapi/model/dto"
)

func ErrorHandlexr(ctx *fiber.Ctx, err error) error {
	fmt.Println("Error Handling is Work")

	code := fiber.StatusInternalServerError
	message := utils.StatusMessage(code)

	if len(err.Error()) > 0 {
		message = err.Error()
	}

	ex, ok := err.(*fiber.Error)
	if ok {
		code = fiber.StatusNotFound
		message = ex.Error()
	}

	_, ok = err.(ValidationError)
	if ok {
		return ctx.JSON(dto.ApiResponse{
			Code:    400,
			Data:    err.Error(),
			Message: message,
		})
	}

	return ctx.JSON(dto.ApiResponse{
		Code:    code,
		Data:    err.Error(),
		Message: message,
	})
}
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	fmt.Println("Error Handling is Work")

	code := fiber.StatusInternalServerError
	message := utils.StatusMessage(code)

	ve, ok := err.(ValidationError)
	if ok {
		code = fiber.StatusBadRequest
		return ctx.Status(code).JSON(dto.ApiResponse{
			Code:    code,
			Message: utils.StatusMessage(code),
			Data: map[string]interface{}{
				"errors": ve.Errors,
			},
		})
	}

	nfe, ok := err.(NotFoundError)
	if ok {
		code = fiber.StatusNotFound
		message = utils.StatusMessage(code)
		if nfe.Message != "" {
			message = nfe.Message
		}
		return ctx.Status(code).JSON(dto.ApiResponse{
			Code:    code,
			Message: message,
		})
	}

	ee, ok := err.(ExceptionError)
	if ok {
		if ee.Code > 0 {
			code = ee.Code
			message = utils.StatusMessage(code)
		}
		if ee.Message != "" {
			message = ee.Message
		}
		return ctx.Status(code).JSON(dto.ApiResponse{
			Code:    code,
			Message: message,
			Data:    ee.Data,
		})
	}

	if len(err.Error()) > 0 {
		message = err.Error()
	}

	var ex *fiber.Error
	ex, ok = err.(*fiber.Error)
	if ok {
		code = ex.Code
		message = ex.Error()
	}

	return ctx.Status(code).JSON(dto.ApiResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
