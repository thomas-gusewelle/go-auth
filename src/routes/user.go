package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thomas-gusewelle/go-auth/src/handlers"
	"github.com/thomas-gusewelle/go-auth/src/middleware"
)

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/user", middleware.WithAuth)
	userGroup.Get("all", handlers.GetAllUsers)
}
