package routes

import (
	"github.com/gofiber/fiber/v2"
	"m3gaplazma/gofiber-restapi/controller"
	"m3gaplazma/gofiber-restapi/middleware"
)

func SetupUserRoutes(router fiber.Router, controller controller.UserController) {
	router.Post("/login", middleware.ApiKeyMiddleware, controller.Login)
	router.Post("/register", middleware.ApiKeyMiddleware, controller.RegisterUser)
	categoryRoutes := router.Group("/users", middleware.ApiKeyMiddleware, middleware.AuthMiddleware)
	categoryRoutes.Get("/", controller.FindAllUsers)
	categoryRoutes.Get("/:userId", controller.FindUserById)
	categoryRoutes.Post("/:userId/activate", controller.ActivateUser)
	categoryRoutes.Post("/:userId/change_password", controller.ChangeUserPassword)
}
