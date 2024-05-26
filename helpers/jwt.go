package helpers

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Helpers struct{
	SecretKey string
}

func NewHelper() HelperInterface {
	return &Helpers{}
}

// Claims structure to hold JWT claims
type Claims struct {
	Role string
	jwt.StandardClaims
}

func (h *Helpers) GenerateToken(userID, role string) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"sub": userID,
		"role": role,
		"exp": time.Now().Add(time.Hour * 8).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

// ValidateJWT validates the JWT token
func (h *Helpers) ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	key := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("token validation failed: %v", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// Convert custom claims to jwt.MapClaims
		mapClaims := jwt.MapClaims{
			"role":     claims.Role,
			"sub":      claims.Subject,
			"exp":      claims.ExpiresAt,
			"iat":      claims.IssuedAt,
		}
		return mapClaims, nil
	}

	return nil, errors.New("invalid token")
}
