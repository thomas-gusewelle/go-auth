package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thomas-gusewelle/go-auth/src/handlers"
)

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/user")
	userGroup.Get("/all", handlers.GetAllUsers)
}
