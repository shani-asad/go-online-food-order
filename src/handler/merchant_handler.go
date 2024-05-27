package handler

import (
	"fmt"
	"online-food/model/dto"
	"online-food/src/usecase"
	"regexp"

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
	}

	c.JSON(201, response)
}

func (h *MerchantHandler) GetMerchants(c *gin.Context) {
	c.JSON(200, gin.H{
		"tes": "",
	})
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
