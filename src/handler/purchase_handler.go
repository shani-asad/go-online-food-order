package handler

import (
	"log"
	"online-food/model/dto"
	"online-food/src/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PurchaseHandler struct {
	iPurchaseUsecase usecase.PurchaseUsecaseInterface
}

func NewPurchaseHandler(iPurchaseUsecase usecase.PurchaseUsecaseInterface) PurchaseHandlerInterface {
	return &PurchaseHandler{iPurchaseUsecase}
}

func (h *PurchaseHandler) GetNearbyMerchants(c *gin.Context) {
	var param dto.RequestNearbyMerchants

	err := c.ShouldBind(&param)
	if err != nil {
		log.Println("Merchant bad request (ShouldBindJSON) >> ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}
	log.Println("param", param)


	validate := validator.New()
	err = validate.Struct(param)
	if err != nil {
		// Validation failed
		for _, err := range err.(validator.ValidationErrors) {
			log.Printf("Error: Field '%s' failed on the '%s' tag\n", err.StructField(), err.Tag())
		}
		c.JSON(400, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	latStr := c.Param("lat")
	longStr := c.Param("long")

	// Convert lat and long from strings to floats
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid latitude"})
		return
	}

	long, err := strconv.ParseFloat(longStr, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid longitude"})
		return
	}
	defaultLimit := 5
	defaultoffset := 0
	if(param.Limit == nil) {param.Limit = &defaultLimit}
	if(param.Offset == nil) {param.Offset = &defaultoffset}

	res, err := h.iPurchaseUsecase.GetNearbyMerchants(long, lat, param)
	if err != nil {
		c.JSON(500, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, res)
}