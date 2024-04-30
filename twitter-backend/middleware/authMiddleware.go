package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/raj5036/twitter-2024/api"
)

var secretKey = []byte("secret-key")

func Authorize(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Running before handler")
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			fmt.Println("Missing Authorization header")
			api.ResponseError(w, "Unauthorized request", http.StatusUnauthorized)
			return
		}

		tokenString = tokenString[len("Bearer "):]
		err := verifyToken(tokenString)
		if err != nil {
			fmt.Println("Invalid Token")
			api.ResponseError(w, "Invalid Authorization Token", http.StatusUnauthorized)
			return
		}

		originalHandler.ServeHTTP(w, r)
		fmt.Println("Running after handler")
	})
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
