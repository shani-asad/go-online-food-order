package handler

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/docker/distribution/uuid"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ImageHandler struct {}

func NewImageHandler() ImageHandlerInterface {
	return &ImageHandler{}
}

const (
	maxUploadSize = 2 * 1024 * 1024 // 2MB
	minUploadSize = 10 * 1024       // 10KB
	
)

var s3Client *s3.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
		return
	}

	// Load the AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create an S3 client
	s3Client = s3.NewFromConfig(cfg)
}

func (h *ImageHandler) UploadImage(c *gin.Context) {
	err := c.Request.ParseMultipartForm(maxUploadSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse form"})
		return
	}

	file, handler, err := c.Request.FormFile("file")
	if file == nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}
	defer file.Close()

	// Validate file size
	if handler.Size > maxUploadSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds 2MB"})
		return
	}
	if handler.Size < minUploadSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size is less than 10KB"})
		return
	}

	// Validate file type
	if !strings.HasSuffix(handler.Filename, ".jpg") && !strings.HasSuffix(handler.Filename, ".jpeg") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type, only .jpg and .jpeg are allowed"})
		return
	}

	// Generate a unique file name
	fileName := fmt.Sprintf("%s%s", uuid.Generate().String(), ".jpeg")

	// Upload the file to S3
	err = uploadFileToS3(file, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	// Respond with the URL of the uploaded file
	imageURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", os.Getenv("AWS_S3_BUCKET_NAME"), fileName)
	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"data": gin.H{
			"imageUrl": imageURL,
		},
	})	
}

func uploadFileToS3(file multipart.File, fileName string) error {
	// Read the file content into a buffer
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}

	// Upload the file to S3
	_, err := s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET_NAME")),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(buf.Bytes()),
		ACL:    "public-read",
	})
	return err
}
