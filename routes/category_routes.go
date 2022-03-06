package routes

import (
	"github.com/gofiber/fiber/v2"
	"m3gaplazma/gofiber-restapi/controller"
)

func SetupCategoryRoutes(router fiber.Router, controller controller.CategoryController) {
	categoryRoutes := router.Group("/categories")
	categoryRoutes.Get("/", controller.FindAllCategories)
	categoryRoutes.Get("/:categoryId", controller.FindCategoryById)
	categoryRoutes.Post("/", controller.CreateCategory)
	categoryRoutes.Put("/:categoryId", controller.UpdateCategory)
	categoryRoutes.Delete("/:categoryId", controller.DeleteCategory)
}
