package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userIdKey contextKey = "userId"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			errResponse(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			errResponse(w, "Authorization header format must be Bearer {token}", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
			}

			return []byte("my-super-sign"), nil
		})

		if err != nil || !token.Valid {
			errResponse(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			errResponse(w, "Invalid claims", http.StatusUnauthorized)
			return
		}

		id, ok := claims["id"].(float64)

		if !ok {
			errResponse(w, "Wrong id type", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), userIdKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}