package routes

import "github.com/gofiber/fiber/v2"

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/user")
	userGroup.Get("/all", userHandler.GetAllUsers)
}
