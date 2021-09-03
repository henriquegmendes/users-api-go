package utils

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"henrique.mendes/users-api/models"
)

func GenerateUserJwt(user models.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	return claims.SignedString([]byte("secret"))
}
