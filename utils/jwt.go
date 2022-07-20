package utils

import (
	"bytes"
	"log"
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

func ParseToken(tokenString string) string {
	_, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (user interface{}, err error) {
		return secretKey, nil
	})

	if err != nil {
		panic("Unable to verify claims")
	}

	out := new(bytes.Buffer)

	// if claims, ok := t.Claims.(*Claims); ok && token.Valid {
	// 	fmt.Fprintf(out, "%v %v", claims.userId, claims.Issuer)
	// }

	return out.String()
}
