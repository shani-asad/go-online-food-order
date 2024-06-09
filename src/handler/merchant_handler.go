package handler

import (
	"fmt"
	"log"
	"online-food/model/dto"
	"online-food/src/usecase"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type MerchantHandler struct {
	iMerchantUsecase usecase.MerchantUsecaseInterface
}

func NewMerchantHandler(iMerchantUsecase usecase.MerchantUsecaseInterface) MerchantHandlerInterface {
	return &MerchantHandler{iMerchantUsecase}
}

func (h *MerchantHandler) CreateMerchant(c *gin.Context) {
	var request dto.RequestCreateMerchant

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	validate := validator.New()
	if err := validate.RegisterValidation("completeURL", validateCompleteURL); err != nil {
		fmt.Println("Failed to register validation function:", err)
		return
	}

	err = validate.Struct(request)
	if err != nil {

		// Validation failed
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Error: Field '%s' failed on the '%s' tag\n", err.StructField(), err.Tag())
		}
		c.JSON(400, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	response, err := h.iMerchantUsecase.CreateMerchant(request)
	if err != nil {
		c.JSON(500, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(201, response)
}

func (h *MerchantHandler) GetMerchants(c *gin.Context) {
	var param dto.RequestGetMerchant

	err := c.ShouldBind(&param)
	if err != nil {
		log.Println("Merchant bad request (ShouldBindJSON) >> ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	res, err := h.iMerchantUsecase.GetMerchants(param)
	if err != nil {
		c.JSON(500, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (h *MerchantHandler) CreateMerchantItem(c *gin.Context) {
	var request dto.RequestCreateMerchantItem

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	validate := validator.New()
	if err := validate.RegisterValidation("completeURL", validateCompleteURL); err != nil {
		fmt.Println("Failed to register validation function:", err)
		return
	}

	err = validate.Struct(request)
	if err != nil {

		// Validation failed
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Error: Field '%s' failed on the '%s' tag\n", err.StructField(), err.Tag())
		}
		c.JSON(400, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	var id dto.RequestBindUrlID
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	merchantID := strconv.Itoa(id.ID)
	merchants, err := h.iMerchantUsecase.GetMerchants(dto.RequestGetMerchant{
		MerchantID: &merchantID,
	})

	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "internal server error"})
		return
	}

	if len(merchants.Data) == 0 {
		c.JSON(404, gin.H{"status": "error", "message": "not found"})
		return
	}

	request.MerchantID = id.ID

	log.Printf("[[[ Create Merchant Item ]]] >>> %+v", request)

	response, err := h.iMerchantUsecase.CreateMerchantItem(request)
	if err != nil {
		c.JSON(500, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(201, response)
}

func (h *MerchantHandler) GetMerchantItems(c *gin.Context) {
	var param dto.RequestGetMerchantItems

	err := c.ShouldBind(&param)
	if err != nil {
		log.Println("Merchant bad request (ShouldBindJSON) >> ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	validate := validator.New()

	err = validate.Struct(param)
	if err != nil {
		// Validation failed
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Error: Field '%s' failed on the '%s' tag\n", err.StructField(), err.Tag())
		}
		c.JSON(400, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	var id dto.RequestBindUrlID
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	merchantID := strconv.Itoa(id.ID)
	merchants, err := h.iMerchantUsecase.GetMerchants(dto.RequestGetMerchant{
		MerchantID: &merchantID,
	})

	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "internal server error"})
		return
	}

	if len(merchants.Data) == 0 {
		c.JSON(404, gin.H{"status": "error", "message": "not found"})
		return
	}

	res, err := h.iMerchantUsecase.GetMerchantItems(param)
	if err != nil {
		c.JSON(500, dto.ResponseStatusAndMessage{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func validateCompleteURL(fl validator.FieldLevel) bool {
	urlString := fl.Field().String()
	// Regular expression to match a complete URL with scheme and valid host
	pattern := `^(http|https)://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}/?`
	matched, err := regexp.MatchString(pattern, urlString)
	if err != nil {
		return false
	}
	return matched
}
