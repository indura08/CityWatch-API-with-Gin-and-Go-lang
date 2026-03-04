package utils

import (
	"citywatch/internal/enums"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func JwtTokenGenerator(email string, role enums.Role, userId int) string {
	claims := jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return ""
	}
	return tokenString
}
