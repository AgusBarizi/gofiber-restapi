package middleware

import (
	"github.com/gofiber/fiber/v2"
	"m3gaplazma/gofiber-restapi/helper"
	"m3gaplazma/gofiber-restapi/model/dto"
)

func ApiKeyMiddleware(ctx *fiber.Ctx) error {
	apiKey := ctx.Get("X-API-Key")
	if apiKey != "secret" {
		return helper.ErrorResponse(ctx, dto.ApiResponse{Code: 401})
	}
	return ctx.Next()
}
