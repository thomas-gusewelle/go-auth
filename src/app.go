package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/thomas-gusewelle/go-auth/db"
	"github.com/thomas-gusewelle/go-auth/src/routes"
	"log"
)

func Setup() error {

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	db.Connect()
	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))

	return nil
}
