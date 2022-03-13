package routes

import (
	"github.com/gofiber/fiber/v2"
	"m3gaplazma/gofiber-restapi/controller"
	"m3gaplazma/gofiber-restapi/middleware"
)

func SetupUserRoutes(router fiber.Router, controller controller.UserController) {
	router.Post("/login", controller.Login)
	router.Post("/register", controller.RegisterUser)
	categoryRoutes := router.Group("/users", middleware.ApiKeyMiddleware)
	categoryRoutes.Get("/", controller.FindAllUsers)
	categoryRoutes.Get("/:userId", controller.FindUserById)
	categoryRoutes.Post("/:userId/activate", controller.ActivateUser)
	categoryRoutes.Post("/:userId/change_password", controller.ChangeUserPassword)
}
