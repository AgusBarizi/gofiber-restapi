package routes

import (
	"github.com/gofiber/fiber/v2"
	"m3gaplazma/gofiber-restapi/config"
	"m3gaplazma/gofiber-restapi/controller"
	"m3gaplazma/gofiber-restapi/repository"
	"m3gaplazma/gofiber-restapi/service"
)

func SetupRoutes(router *fiber.App) {
	api := router.Group("/api")

	categoryRepositoryImpl := repository.NewCategoryRepository(config.DB)
	categoryServiceImpl := service.NewCategoryService(categoryRepositoryImpl)
	categoryControllerImpl := controller.NewCategoryController(categoryServiceImpl)
	SetupCategoryRoutes(api, categoryControllerImpl)

	productRepositoryImpl := repository.NewProductRepository(config.DB)
	productServiceImpl := service.NewProductService(productRepositoryImpl)
	productControllerImpl := controller.NewProductController(productServiceImpl)
	SetupProductRoutes(api, productControllerImpl)
}
