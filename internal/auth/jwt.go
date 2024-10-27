package auth

import (
	"ecommerce-cloning-app/internal/logger"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type claims struct {
	Phone string
	jwt.RegisteredClaims
}

func GenerateJWT(phone string) (string, error) {
	expirationTime := time.Now().Local().Add(time.Minute * 30)
	claims := &claims{
		Phone: phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func GenerateRefreshToken() string {
	return uuid.NewString()
}

func ValidateJWT(tokenString string) (*claims, error) {
	claims := &claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	if time.Now().Local().After(claims.ExpiresAt.Time) {
		logger.Logging().Error("token has expired")
		return nil, err
	}

	return claims, nil
}
