package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thomas-gusewelle/go-auth/src/handlers"
)

func IndexRoutes(app *fiber.App) {
	index := app.Group("/")
	index.Get("/", handlers.Index)
	index.Get("/signin", handlers.SignIn)
	index.Get("/signup", handlers.SignUp)
}
