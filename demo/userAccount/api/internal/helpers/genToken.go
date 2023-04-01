package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenToken(name string, secret_key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"exp":  time.Now().Add(time.Second * 600).Unix(),
	})
	return token.SignedString([]byte(secret_key))
}
