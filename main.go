package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	log "github.com/sirupsen/logrus"
	"m3gaplazma/gofiber-restapi/config"
	"m3gaplazma/gofiber-restapi/exception"
	"m3gaplazma/gofiber-restapi/routes"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//log.SetLevel(log.WarnLevel)
}

func main() {
	app := fiber.New(config.NewFiberConfig())
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${method} ${path} ${status} ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/jakarta",
	}))
	app.Static("/storage", "./storage")

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
