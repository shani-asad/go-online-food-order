package helpers

import "github.com/golang-jwt/jwt"

type AuthHelperInterface interface {
	GenerateToken(userId, role string) (string, error)
	ValidateJWT(tokenString string) (jwt.MapClaims, error)
}
