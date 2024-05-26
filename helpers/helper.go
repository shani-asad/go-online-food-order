package helpers

import "github.com/golang-jwt/jwt"

type HelperInterface interface {
	GenerateToken(userId, role string) (string, error)
	ValidateJWT(tokenString string) (jwt.MapClaims, error)
}
