package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(id, role uint, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add((7 * 24 * time.Hour)).Unix()
	claims["role"] = role
	claims["id"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
