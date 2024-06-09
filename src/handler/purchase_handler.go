package handler

import (
	"log"
	"online-food/model/dto"
	"online-food/src/usecase"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PurchaseHandler struct {
	iPurchaseUsecase usecase.PurchaseUsecaseInterface
	iMerchantUsecase usecase.MerchantUsecaseInterface
}

func NewPurchaseHandler(iPurchaseUsecase usecase.PurchaseUsecaseInterface, iMerchantUsecase usecase.MerchantUsecaseInterface) PurchaseHandlerInterface {
	return &PurchaseHandler{iPurchaseUsecase, iMerchantUsecase}
}

func (h *PurchaseHandler) GetNearbyMerchants(c *gin.Context) {
	var param dto.RequestNearbyMerchants

	err := c.ShouldBind(&param)
	if err != nil {
		log.Println("Merchant bad request (ShouldBind) >> ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(param)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Printf("Error: Field '%s' failed on the '%s' tag\n", err.StructField(), err.Tag())
		}
		c.JSON(400, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	latLong := c.Param("latLong")

	arr := strings.Split(latLong, ",")
	latStr := arr[0]
	longStr := arr[1]

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

func (h *PurchaseHandler) CreateEstimation(c *gin.Context) {
	var param dto.RequestEstimate
	err := c.ShouldBindJSON(&param)
	if err != nil {
		log.Println("Merchant bad request (ShouldBindJSON) >> ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(param)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Printf("Error: Field '%s' failed on the '%s' tag\n", err.StructField(), err.Tag())
		}
		c.JSON(400, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	countIsStartingPoint := 0
	for _, v := range param.Orders {
		if(v.IsStartingPoint) {countIsStartingPoint++}
	}
	if(countIsStartingPoint != 1) {
		c.JSON(400, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: "There must be exactly one order with IsStartingPoint set to true",
		})
		return
	}

	merchantIds := []string{}
	itemIds := []string{}
	for _, o := range param.Orders {
		merchantIds = append(merchantIds, o.MerchantId)
		for _, i := range o.Items {
			itemIds = append(itemIds, i.ItemId)
		}
	}

	merchantIds = UniqueStrings(merchantIds)
	itemIds = UniqueStrings(itemIds)

	merchantIdsString := strings.Join(merchantIds, ",")
	merchantCount := h.iMerchantUsecase.GetMerchantCountByIds(merchantIdsString)

	if merchantCount != len(merchantIds) {
		c.JSON(404, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: "Some merchant id does not exist",
		})
		return
	}
	
	itemIdsString := strings.Join(itemIds, ",")
	itemCount := h.iMerchantUsecase.GetItemCountByIds(itemIdsString)

	if itemCount != len(itemIds) {
		c.JSON(404, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: "Some item id does not exist",
		})
		return
	}

	userId, exist := c.Get("user_id")
	if(!exist){
		c.JSON(404, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: "failed c.Get(\"user_id\")",
		})
		return
	}
	res, err := h.iPurchaseUsecase.CreateEstimation(param, userId.(string))

	if res.EstimatedDeliveryTimeInMinutes == -1 {
		c.JSON(400, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(500, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, res)

}

func (h *PurchaseHandler) CreateOrder(c *gin.Context) {
	var param dto.RequestOrder

	err := c.ShouldBindJSON(&param)
	if err != nil {
		log.Println("Merchant bad request (ShouldBindJSON) >> ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(param)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Printf("Error: Field '%s' failed on the '%s' tag\n", err.StructField(), err.Tag())
		}
		c.JSON(400, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	// orderId, err := h.iPurchaseUsecase.CreateOrder(strconv.Itoa(param.CalculatedEstimateId))
	orderId, err := h.iPurchaseUsecase.CreateOrder(param.CalculatedEstimateId)
	
	if(orderId == ""){
		c.JSON(400, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: "estimate ID not found",
		})
		return
	}
	
	if err != nil {
		c.JSON(500, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(201, dto.ResponseCreateOrder{
		OrderId: orderId,
	})
}

func (h *PurchaseHandler) GetOrders(c *gin.Context) {
	var param dto.RequestGetOrders
	err := c.ShouldBind(&param)
	if err != nil {
		log.Println("Merchant bad request (ShouldBindJSON) >> ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	log.Println("paramparamparamparam", param)

	validate := validator.New()
	err = validate.Struct(param)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Printf("Error: Field '%s' failed on the '%s' tag\n", err.StructField(), err.Tag())
		}
		c.JSON(400, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	defaultLimit := 5
	defaultoffset := 0
	if(param.Limit == nil) {param.Limit = &defaultLimit}
	if(param.Offset == nil) {param.Offset = &defaultoffset}

	userId, exist := c.Get("user_id")

	if !exist {
		c.JSON(401, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: "Cannot get userId",
		})
		return
	}

	log.Println("userId>>>", userId)

	res, _ := h.iPurchaseUsecase.GetOrders(param, userId.(string))

	c.JSON(200, res)
}

func UniqueStrings(input []string) []string {
    // Create a map to track unique values.
    uniqueMap := make(map[string]struct{})
    
    // Iterate over the input slice and add each unique value to the map.
    for _, str := range input {
        uniqueMap[str] = struct{}{}
    }

    // Create a slice to hold the distinct values.
    uniqueSlice := make([]string, 0, len(uniqueMap))
    
    // Populate the slice with the keys from the map.
    for key := range uniqueMap {
        uniqueSlice = append(uniqueSlice, key)
    }
    
    return uniqueSlice
}