package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thomas-gusewelle/go-auth/db"
)

var sampleKey = "HelloWordKey"

func GenerateJWT(user *db.User) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = user.Name

	jwtString, err := token.SignedString(sampleKey)
	if err != nil {
		return "", err
	}
	return jwtString, err
}
