package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router *fiber.App) {
	api := router.Group("/api")
	SetupProductRoutes(api)
	//routes.Get("/api/category/:categoryId", categoryController.FindById)
	//routes.Post("/api/category", categoryController.Create)
	//routes.Put("/api/category/:categoryId", categoryController.Update)
	//routes.Delete("/api/category/:categoryId", categoryController.Delete)
}
