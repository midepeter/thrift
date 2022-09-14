package middleware

import (
	"net/http"

	"github.com/midepeter/thrift/pkg/jwt"
)

//Authorization middleware
func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Check the request header here
		if r.URL.Path == "/user.User/Register" || r.URL.Path == "/user.User/SignIn" {
			return
		}

		tokenString := r.Header.Get("token")
		//validate token
		_, _, err := jwt.ParseToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid token"))
			return
		}

		h.ServeHTTP(w, r)
	})
}
