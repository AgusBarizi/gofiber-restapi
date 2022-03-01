package config

import (
	"github.com/gofiber/fiber/v2"
	"m3gaplazma/gofiber-restapi/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		GETOnly:      false,
		ErrorHandler: exception.ErrorHandler,
	}
}
