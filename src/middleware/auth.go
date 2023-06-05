package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/thomas-gusewelle/go-auth/auth"
)

func WithAuth(c *fiber.Ctx) error {
	jtoken := c.Cookies("jwt")
	fmt.Println("MIddleware token", jtoken)
	verifiedToken, err := auth.VerfyJWT(jtoken)
	if err != nil {
		fmt.Println("Middleware error: ", err)
		return c.Redirect("/signin")
	}
	fmt.Println(verifiedToken)
	return c.Next()
}
