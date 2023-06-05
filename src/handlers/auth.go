package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/thomas-gusewelle/go-auth/auth"
	"github.com/thomas-gusewelle/go-auth/db"
	"github.com/thomas-gusewelle/go-auth/utils"
	"golang.org/x/crypto/bcrypt"
)

func PostSignUp(c *fiber.Ctx) error {
	var user db.User
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	encPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	utils.CheckError(err)

	user.Name = name
	user.Email = email
	password = string(encPass)

	db.Database.Create(&user)

	jwt, err := auth.GenerateJWT(&user)
	fmt.Println("JWT Print: ", jwt)
	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = jwt
	c.Cookie(cookie)
	return c.Redirect("/user/all")
}

func PostSignIn(c *fiber.Ctx) error {

	var user db.User
	email := c.FormValue("email")
	password := c.FormValue("password")
	db.Database.Where("email = ?", email).First(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return c.Redirect("/users")
}

func PostSignOut(c *fiber.Ctx) error {
	c.ClearCookie("jwt")
	return c.Redirect("/")
}
