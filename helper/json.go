package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"m3gaplazma/gofiber-restapi/model/dto"
)

func SuccessResponser(ctx *fiber.Ctx, data interface{}, message ...string) error {
	code := fiber.StatusOK

	apiResponse := dto.ApiResponse{
		Code:    code,
		Data:    data,
		Message: utils.StatusMessage(code),
	}

	if len(message) > 0 {
		apiResponse.Message = message[0]
	}
	return ctx.Status(code).JSON(apiResponse)
}

func ErrorResponser(ctx *fiber.Ctx, message string, data interface{}, httpCode ...int) error {
	code := fiber.StatusInternalServerError
	if len(httpCode) > 0 {
		code = httpCode[0]
	}

	apiResponse := dto.ApiResponse{
		Code:    code,
		Data:    data,
		Message: message,
	}
	return ctx.Status(code).JSON(apiResponse)
}

func SuccessResponse(ctx *fiber.Ctx, response dto.ApiResponse) error {
	if response.Code == 0 {
		response.Code = fiber.StatusOK
	}

	if response.Message == "" {
		response.Message = utils.StatusMessage(response.Code)
	}
	return ctx.Status(response.Code).JSON(response)
}

func ErrorResponse(ctx *fiber.Ctx, response dto.ApiResponse) error {
	if response.Code == 0 {
		response.Code = fiber.StatusInternalServerError
	}
	if response.Message == "" {
		response.Message = utils.StatusMessage(response.Code)
	}
	return ctx.Status(response.Code).JSON(response)
}
