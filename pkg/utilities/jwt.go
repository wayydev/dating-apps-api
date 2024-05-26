package utilities

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte(os.Getenv("APP_JWT_KEY"))

type JWT struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func CreateTokenJWT(request *JWT) (string, error) {
	if len(JwtKey) == 0 {
		return "", fmt.Errorf("JWT Key Not Found")
	}

	expiresAt := time.Now().Add(24 * time.Hour)
	claims := JWT{
		request.ID,
		request.Name,
		request.Username,
		request.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	generate := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := generate.SignedString(JwtKey)
	if err != nil {
		return token, err
	}

	return token, nil
}

func ParseTokenJWT(token string) (*JWT, error) {
	parse, err := jwt.ParseWithClaims(token, &JWT{}, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	return parse.Claims.(*JWT), nil
}
