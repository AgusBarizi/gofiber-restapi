package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"m3gaplazma/gofiber-restapi/config"
	"m3gaplazma/gofiber-restapi/exception"
	"m3gaplazma/gofiber-restapi/model/domain"
	"m3gaplazma/gofiber-restapi/routes"
)

var product domain.Product

//func getSingleProduct(ctx *fiber.Ctx) error {
//	//product := domain.Product{
//	//	Id:    0,
//	//	Name:  "Apple Macbook Pro M1",
//	//	Sku:   "AM2021",
//	//	Stock: 2,
//	//}
//	return ctx.Status(fiber.StatusOK).JSON(product)
//}

func createProduct(ctx *fiber.Ctx) error {
	body := new(domain.Product)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}

	product = domain.Product{
		Id:    1,
		Sku:   body.Sku,
		Stock: body.Stock,
		Name:  body.Name,
	}

	return ctx.Status(fiber.StatusOK).JSON(product)
}

func main() {
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	//app.Get("/api", func(ctx *fiber.Ctx) error {
	//	return ctx.SendString("GoFiber RestApi")
	//})

	config.ConnectSQL()
	routes.SetupRoutes(app)

	err := app.Listen(
		fmt.Sprintf("%s:%s",
			config.Env("APP_HOST"),
			config.Env("APP_PORT"),
		))
	exception.PanicIfError(err)
}
