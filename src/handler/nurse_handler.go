package handler

import (
	"errors"
	"fmt"
	"health-record/model/dto"
	"health-record/src/usecase"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type NurseHandler struct {
	iNurseUsecase usecase.NurseUsecaseInterface
}

func NewNurseHandler(iNurseUsecase usecase.NurseUsecaseInterface) NurseHandlerInterface {
	return &NurseHandler{iNurseUsecase}
}


func (h *NurseHandler) RegisterNurse(c *gin.Context) {
	var request dto.RequestCreateNurse
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Register bad request")
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	// Validate request payload
	err = ValidateRegisterNurseRequest(request.Nip, request.Name)
	if err != nil {
		log.Println("Register bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	if request.IdentityCardScanImg == "" {
		log.Println("Register bad request > invalid IdentityCardScanImg")
		c.JSON(400, gin.H{"status": "bad request", "message": "invalid IdentityCardScanImg"})
		return
	}

	// Check if email already exists
	exists, _ := h.iNurseUsecase.GetNurseByNIP(request.Nip)
	if exists {
		log.Println("Register bad request ", err)
		c.JSON(409, gin.H{"status": "bad request", "message": "nip already exists"})
		return
	}

	userId, err := h.iNurseUsecase.RegisterNurse(request)
	if err != nil {
		log.Println("Register bad request ", err)
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	log.Println("Register successful")
	c.JSON(201, gin.H{
    "message": "Nurse registered successfully",
    "data": gin.H{
			"userId": userId,
			"nip": request.Nip, 
			"name": request.Name,
		},
	})
}

func (h *NurseHandler) GetUsers(c *gin.Context) {
	var userId, nip, name, role string
	
	if _, ok := c.Request.URL.Query()["userId"]; ok{
		userId = c.Query("userId")
	}

	if _, ok := c.Request.URL.Query()["nip"]; ok{
		nip = c.Query("nip")
	}

	if _, ok := c.Request.URL.Query()["name"]; ok{
		name = c.Query("name")
	}

	if _, ok := c.Request.URL.Query()["role"]; ok{
		role = c.Query("role")
	}

	var limit, offset int
	if _, ok := c.Request.URL.Query()["limit"]; ok && c.Query("limit") != "" {
		val, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
				limit = 5
		} else {
			limit = val
		}
	} else {
		limit = 5
	}

	if _, ok := c.Request.URL.Query()["offset"]; ok && c.Query("offset") != "" {
		val, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
				offset = 0
		} else {
			offset = val
		}
	}

	var createdAt  string
	
	if _, ok := c.Request.URL.Query()["createdAt"]; ok && c.Query("createdAt") != "" {
		createdAt  = c.Query("createdAt")
	}

	params := dto.RequestGetUser{
		UserId	: userId,
		Limit	: limit,
		Offset	: offset,
		Name: name,
		NIP: nip,
		Role: role,
		CreatedAt	: createdAt,
	}

	fmt.Println("paramsCustomerId>>>>>", params.UserId)
	fmt.Println("paramsCustomerId>>>>>", params.NIP)
	fmt.Println("paramsCustomerId>>>>>", params.Role)
	fmt.Println("paramsCustomerId>>>>>", params.Name)
	fmt.Println("paramsLimit>>>>>", params.Limit)
	fmt.Println("paramsOffset>>>>>", params.Offset)
	fmt.Println("paramsCreatedAt>>>>>", params.CreatedAt)

	users, err := h.iNurseUsecase.GetUsers(params)

	if err != nil {
		log.Println("get sku server error ", err)
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	if(len(users) < 1) { users = []dto.UserDTO{}}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": users})
}

func (h *NurseHandler) UpdateNurse(c *gin.Context) {
	userId := c.Param("userId")
	var request dto.RequestUpdateNurse
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Update bad request")
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	user, err := h.iNurseUsecase.GetNurseByID(userId)
	if err != nil {
		log.Println("Update bad request ", err)
		c.JSON(404, gin.H{"status": "bad request", "message": "userId not found"})
	}

	nStr := strconv.FormatInt(user.Nip, 10)
	if !strings.HasPrefix(nStr, "303") {
		c.JSON(404, gin.H{"status": "bad request", "message": "user not found"})
			return
	}
	exists, _ := h.iNurseUsecase.GetNurseByNIP(request.Nip)
	if exists {
		log.Println("Update bad request ", err)
		c.JSON(409, gin.H{"status": "bad request", "message": "nip already exists"})
		return
	}

	// Validate request payload
	err = ValidateRegisterNurseRequest(request.Nip, request.Name)
	if err != nil {
		log.Println("Update bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}
	
	statusCode := h.iNurseUsecase.UpdateNurse(userId, request)

	c.JSON(statusCode, gin.H{"status": statusCode})
}

func (h *NurseHandler) DeleteNurse(c *gin.Context) {
	userId := c.Param("userId")
	user, err := h.iNurseUsecase.GetNurseByID(userId)
	if err != nil {
		log.Println("Update bad request ", err)
		c.JSON(404, gin.H{"status": "bad request", "message": "userId not found"})
	}
	nStr := strconv.FormatInt(user.Nip, 10)
	if !strings.HasPrefix(nStr, "303") {
		c.JSON(404, gin.H{"status": "bad request", "message": "user not found"})
			return
	}
	statusCode := h.iNurseUsecase.DeleteNurse(userId)

	c.JSON(statusCode, gin.H{"status": statusCode})
}

func (h *NurseHandler) AddAccess(c *gin.Context) {
	userId := c.Param("userId")
	var request dto.RequestAddAccess
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("add access bad request")
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}
	// Validate request payload
	err = validatePassword(request.Password)
	if err != nil {
		log.Println("Update bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}
	statusCode := h.iNurseUsecase.AddAccess(userId, request)

	c.JSON(statusCode, gin.H{"status": statusCode})
}

func validatePassword(password string) error {
	if len(password) < 5 || len(password) > 33 {
		return errors.New("password length must be between 5 and 33 characters")
	}

	return nil
}

func ValidateRegisterNurseRequest(nip int64, name string) error {
	// Validate email format
	if !isValidNipNurse(nip) {
		return errors.New("nip must be in valid nip format")
	}

	// Validate name length
	if len(name) < 5 || len(name) > 50 {
		return errors.New("name length must be between 5 and 50 characters")
	}

	return nil
}

func isValidNipNurse(nip int64) bool {
	nipRegex := fmt.Sprintf(`^303[12](200[0-9]|201[0-9]|202[0-4])(0[1-9]|1[0-2])\d{3,5}$`)
	// Convert the nip int64 to a string
	nipStr := strconv.FormatInt(nip, 10)
	// Match the string with the regex
	match, _ := regexp.MatchString(nipRegex, nipStr)
	return match
}