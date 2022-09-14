package jwt

import (
	"fmt"
	"log"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var (
	secretKey = "my_secret_key"
)

func GenerateJwtToken(Email string) (string, error) {

	claims := Claims{
		Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "admin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println("The error", err)
		panic(err)
	}

	return ss, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
	if tokenString == "" {
		return nil, nil, fmt.Errorf("Invalid token string")
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (user interface{}, err error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, nil, fmt.Errorf("Unauthorized token")
	}

	return token, claims, nil
}
