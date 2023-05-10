package handlers

import "github.com/gofiber/fiber/v2"

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func SignIn(c *fiber.Ctx) error {
	return c.Render("signin", fiber.Map{})
}

func SignUp(c *fiber.Ctx) error {
	return c.Render("signup", fiber.Map{})
}

func PostSingUp(c *fiber.Ctx) error {
	return c.Redirect("/users")
}
