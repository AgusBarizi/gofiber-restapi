package routes

import (
	"github.com/gofiber/fiber/v2"
	"m3gaplazma/gofiber-restapi/controller"
	"m3gaplazma/gofiber-restapi/middleware"
)

func SetupProductRoutes(router fiber.Router, controller controller.ProductController) {
	productRoutes := router.Group("/products", middleware.ApiKeyMiddleware, middleware.AuthMiddleware)
	productRoutes.Get("/", controller.FindAllProducts)
	productRoutes.Get("/:productId", controller.FindProductDetail)
	productRoutes.Post("/", controller.CreateProduct)
	productRoutes.Put("/:productId", controller.UpdateProduct)
	productRoutes.Delete("/:productId", controller.DeleteProduct)
}
