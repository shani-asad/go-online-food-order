package handler

import (
	"github.com/gin-gonic/gin"
)

type AuthHandlerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	LoginNurse(c *gin.Context)
}

type NurseHandlerInterface interface {
	RegisterNurse(c *gin.Context)
	UpdateNurse(c *gin.Context)
	DeleteNurse(c *gin.Context)
	GetUsers(c *gin.Context)
	AddAccess(c *gin.Context)
}

type PatientHandlerInterface interface {
	CreatePatient(c *gin.Context)
	GetPatients(c *gin.Context)
	CreateRecord(c *gin.Context)
	GetRecords(c *gin.Context)
}

type ImageHandlerInterface interface {
	UploadImage(c *gin.Context)
}
