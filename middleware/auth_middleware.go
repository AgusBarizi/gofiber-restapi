package middleware

import (
	"github.com/gofiber/fiber/v2"
	"m3gaplazma/gofiber-restapi/helper"
	"m3gaplazma/gofiber-restapi/model/dto"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token == "" {
		return helper.ErrorResponse(ctx, dto.ApiResponse{Code: 401})
	}
	payload, err := helper.DecodeToken(token)
	if err != nil {
		return helper.ErrorResponse(ctx, dto.ApiResponse{Code: 401})
	}
	data := payload["data"].(map[string]interface{})
	ctx.Locals("userInfo", data)
	return ctx.Next()
}
