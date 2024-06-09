package helpers

import "github.com/golang-jwt/jwt"

type AuthHelperInterface interface {
	GenerateToken(userId, role string) (string, error)
	ValidateJWT(tokenString string) (jwt.MapClaims, error)
}

type DistanceHelperInterface interface {
	GetHaversineDistance(lat1, lon1, lat2, lon2 float64) float64
}