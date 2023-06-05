package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thomas-gusewelle/go-auth/src/handlers"
)

func AuthRoutes(app *fiber.App) {
	authGroup := app.Group("/auth")
	authGroup.Post("/signin", handlers.PostSignIn)
	authGroup.Post("/signout", handlers.PostSignOut)
	authGroup.Post("signup", handlers.PostSignUp)
}
