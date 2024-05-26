package handler

import (
	"github.com/gin-gonic/gin"
)

type AuthHandlerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type ImageHandlerInterface interface {
	UploadImage(c *gin.Context)
}

type MerchantHandlerInterface interface {
	CreateMerchant(c *gin.Context)
	GetMerchants(c *gin.Context)
}
