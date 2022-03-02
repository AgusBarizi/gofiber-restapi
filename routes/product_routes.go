package routes

import (
	"github.com/gofiber/fiber/v2"
	"m3gaplazma/gofiber-restapi/controller"
)

func SetupProductRoutes(router fiber.Router, controller controller.ProductController) {
	productRoutes := router.Group("/products")
	productRoutes.Get("/", controller.FindAllProducts)
	productRoutes.Get("/:productId", controller.FindProductById)
	productRoutes.Post("/", controller.CreateProduct)
	productRoutes.Put("/:productId", controller.UpdateProduct)
	productRoutes.Delete("/:productId", controller.DeleteProduct)
}
