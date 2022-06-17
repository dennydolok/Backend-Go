package helper

import (
	"WallE/config"
	"strings"
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

func GetClaim(reqToken string) (int, int) {
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.InitConfig().SECRET_KEY), nil
	})
	if err != nil {
		return 0, 0
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims["id"].(int), claims["role"].(int)
}
