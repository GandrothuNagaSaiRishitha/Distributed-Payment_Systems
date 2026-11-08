package main

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// jwt structure
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

			// Accept only the expected signing algorithm.
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("unexpected signing algorithm")
			}

			return []byte(secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
