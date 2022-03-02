package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"m3gaplazma/gofiber-restapi/config"
	"m3gaplazma/gofiber-restapi/exception"
	"m3gaplazma/gofiber-restapi/routes"
)

func main() {
	app := fiber.New(config.NewFiberConfig())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("GoFiber RestApi")
	})

	app.Use(recover.New())
	config.ConnectSQL()
	routes.SetupRoutes(app)

	err := app.Listen(
		fmt.Sprintf("%s:%s",
			config.Env("APP_HOST"),
			config.Env("APP_PORT"),
		))
	exception.PanicIfError(err)
}
