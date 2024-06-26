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
	CreateMerchantItem(c *gin.Context)
	GetMerchantItems(c *gin.Context)
}

type PurchaseHandlerInterface interface {
	GetNearbyMerchants(c *gin.Context)
	CreateEstimation(c *gin.Context)
	CreateOrder(c *gin.Context)
	GetOrders(c *gin.Context)
}
