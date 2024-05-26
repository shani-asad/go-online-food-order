package handler

import (
	"errors"
	"log"
	"online-food/model/dto"
	"online-food/src/usecase"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	iAuthUsecase usecase.AuthUsecaseInterface
}

func NewAuthHandler(iAuthUsecase usecase.AuthUsecaseInterface) AuthHandlerInterface {
	return &AuthHandler{iAuthUsecase}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request dto.RequestCreateUser
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Register bad request")
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	// Validate request payload
	if(request.Email == "") {
		log.Println("Register bad request > invalid Email")
		c.JSON(400, gin.H{"status": "bad request", "message": "invalid Email"})
		return
	}

	err = ValidateRegisterRequest(request.Email, request.Username, request.Password)
	if err != nil {
		log.Println("Register bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	// Check if email already exists
	exists, _ := h.iAuthUsecase.GetUserByUsername(request.Username)
	if exists {
		log.Println("Register bad request ", err)
		c.JSON(409, gin.H{"status": "bad request", "message": "username already exists"})
		return
	}

	path := c.Request.URL.Path
	role := ""

	if strings.HasPrefix(path, "/admin") {
			role = "admin"
	} else {
			role = "user"
	}

	// Check if email already exists in the same role
	exist, _ := h.iAuthUsecase.GetExistingUserInTheRoleByEmail(request.Email, role)
	if exist {
		log.Println("Register bad request ", err)
		c.JSON(409, gin.H{"status": "bad request", "message": "email already exists"})
		return
	}

	token, err := h.iAuthUsecase.Register(request, role)
	if err != nil {
		log.Println("Register bad request ", err)
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	log.Println("Register successful")
	c.JSON(201, gin.H{
    "token": token,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request dto.RequestAuth
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Login bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}
	
	if(request.Username == "") {
		log.Println("Register bad request > invalid Username")
		c.JSON(400, gin.H{"status": "bad request", "message": "invalid username"})
		return
	}

	err = ValidateLoginRequest(request.Username, request.Password)
	if err != nil {
		log.Println("Login bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	path := c.Request.URL.Path
	role := ""

	if strings.HasPrefix(path, "/admin") {
			role = "admin"
	} else {
			role = "user"
	}

	token, status := h.iAuthUsecase.Login(request, role)
	if status == 400 {
		log.Println("Login bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": "bad request"})
		return
	} else if (status == 404) {
		c.JSON(404, gin.H{"status": "bad request", "message": "user not found"})
		return
	}

	log.Println("Login successful")
	c.JSON(200, gin.H{
    "token": token,
	})
}

// ValidateRegisterRequest validates the register user request payload
func ValidateRegisterRequest(email, username, password string) error {
	// Validate email format
	if !isValidEmail(email) {
		return errors.New("email invalid")
	}

	// Validate name length
	if len(username) < 5 || len(username) > 30 {
		return errors.New("name length must be between 5 and 50 characters")
	}

	// Validate password length
	if len(password) < 5 || len(password) > 30 {
		return errors.New("password length must be between 5 and 15 characters")
	}

	return nil
}

func ValidateLoginRequest(username, password string) error {
	// Validate nip format

	if len(username) < 5 || len(username) > 30 {
		return errors.New("name length must be between 5 and 30 characters")
	}

	// Validate password length
	if len(password) < 5 || len(password) > 30 {
		return errors.New("password length must be between 5 and 30 characters")
	}

	return nil
}

func isValidEmail(email string) bool {
	// Regular expression pattern for email format
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}
