package utils

import (
	"bytes"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateToken(secretKey, userId string) string {

	claims := Claims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "admin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}

	return ss
}

func ParseToken(tokenString string) string {
	t, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (user interface{}, err error) {
		user = []byte("secret_key")
		err = nil
		return
	})

	if err != nil {
		panic("Unable to verify claims")
	}

	out := new(bytes.Buffer)

	if claims, ok := t.Claims.(*Claims); ok && token.Valid {
		fmt.Fprintf(out, "%v %v", claims.userId, claims.Issuer)
	}

	return out.String()
}
