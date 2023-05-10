package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thomas-gusewelle/go-auth/db"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []db.User
	query := db.Database.Find(&users)
	return c.Render("users", fiber.Map{})
}
