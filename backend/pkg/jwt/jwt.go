package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(userID string) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    "react-golang-playground",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		ID:        userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}
