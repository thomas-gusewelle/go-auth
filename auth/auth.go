package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thomas-gusewelle/go-auth/db"
)

var sampleKey = []byte("HelloWordKey")

type Claims struct {
	Username string
	jwt.RegisteredClaims
}

func GenerateJWT(user *db.User) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtString, err := token.SignedString(sampleKey)
	if err != nil {
		fmt.Println("Erorr: ", err)
		return "", err
	}
	return jwtString, err
}

func VerfyJWT(jtoken string) (*jwt.Token, error) {
	// claims := &Claims{}

	// token, err := jwt.ParseWithClaims(jtoken, claims, func(t *jwt.Token) (interface{}, error) {
	//    _, ok := t.Method.(* jwt.Sign)
	// 	return jtoken, nil
	// })
	token, err := jwt.Parse(jtoken, func(t *jwt.Token) (interface{}, error) {
		return sampleKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
