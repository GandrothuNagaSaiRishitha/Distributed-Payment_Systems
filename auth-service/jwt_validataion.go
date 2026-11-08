package main

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func validateJWT(tokenString string) (*Claims, error) {

	secret := os.Getenv("JWT_SECRET")

	if len(secret) < 32 {
		return nil, errors.New("JWT secret is missing or too short")
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithIssuer("auth-service"),
		jwt.WithExpirationRequired(),
	)

	if err != nil {
		return nil, err
	}
	//if the expiry date/time has passed, declare the token as invalid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.Subject == "" {
		return nil, errors.New("token subject is missing")
	}

	return claims, nil
}
