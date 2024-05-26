package middleware

import (
	"health-record/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	helper helpers.HelperInterface
}

type MiddlewareInterface interface {
	AuthMiddleware(c *gin.Context)
	RoleMiddleware(role string) gin.HandlerFunc
}

func NewMiddleware(helper helpers.HelperInterface) MiddlewareInterface {
	return &Middleware{helper}
}

func (m *Middleware) AuthMiddleware(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")

	if authorizationHeader == "" {
		c.JSON(http.StatusUnauthorized, "request token is missing or expired")
		c.Abort()
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", 1)
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, "invalid token format")
		c.Abort()
		return
	}

	// Validate JWT
	claims, err := m.helper.ValidateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "invalid token")
		c.Abort()
		return
	}

	c.Set("user_id", claims["sub"])
	c.Set("user_role", claims["role"])
	c.Next()
}

func (m *Middleware) RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		
		if !exists || role != requiredRole {
			c.JSON(http.StatusBadRequest, "you don't have permission to access this resource")
			c.Abort()
			return
		}
		c.Next()
	}
}
