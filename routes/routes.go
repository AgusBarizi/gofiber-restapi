package routes

import (
	"github.com/gofiber/fiber/v2"
	"m3gaplazma/gofiber-restapi/config"
	"m3gaplazma/gofiber-restapi/controller"
	"m3gaplazma/gofiber-restapi/repository"
)

func SetupRoutes(router *fiber.App) {
	api := router.Group("/api")
	productRepositoryImpl := repository.NewProductRepository(config.DB)
	productControllerImpl := controller.NewProductController(productRepositoryImpl)
	SetupProductRoutes(api, productControllerImpl)
}
