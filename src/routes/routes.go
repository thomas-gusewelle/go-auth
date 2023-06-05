package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	IndexRoutes(app)
	AuthRoutes(app)
	UserRoutes(app)
}
