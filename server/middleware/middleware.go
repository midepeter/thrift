package middleware

import (
	"log"
	"net/http"
)

//Authorization middleware
func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Check the request header here
		log.Println("Okay, It has gotten to the handler")
		h.ServeHTTP(w, r)
	})
}
