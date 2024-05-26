package handler

import (
	"online-food/src/usecase"

	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	iMerchantUsecase usecase.MerchantUsecaseInterface
}

func NewMerchantHandler(iMerchantUsecase usecase.MerchantUsecaseInterface) MerchantHandlerInterface {
	return &MerchantHandler{iMerchantUsecase}
}

func (h *MerchantHandler) CreateMerchant(c *gin.Context) {
	c.JSON(201, gin.H{
		"tes": "",
	})
}

func (h *MerchantHandler) GetMerchants(c *gin.Context) {
	c.JSON(200, gin.H{
		"tes": "",
	})
}
