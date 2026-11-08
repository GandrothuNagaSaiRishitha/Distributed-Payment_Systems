// why this middleware?: security checkpoint
package main

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const userClaimsKey contextKey = "userClaims"

func authMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || parts[1] == "" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		claims, err := validateJWT(parts[1])

		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userClaimsKey, claims)

		next.ServeHTTP(w, r.WithContext(ctx)) //Authentication succeeded. Now pass this request to the next handler.
	})
}
