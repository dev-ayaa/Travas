package token

import "github.com/golang-jwt/jwt/v4"

type TravasClaims struct {
	jwt.RegisteredClaims
	Email string
}
