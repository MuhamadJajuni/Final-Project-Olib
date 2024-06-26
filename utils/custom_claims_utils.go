package utils

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	jwt.RegisteredClaims
	Name  string `json:"Name"`
	Email string `json:"Email"`
}
