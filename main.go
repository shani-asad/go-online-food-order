package main

import (
	"fmt"
	"log"
	"online-food/app/server"
	"online-food/db"
	"online-food/helpers"
	"online-food/model/properties"
	"online-food/src/handler"
	"online-food/src/middleware"
	"online-food/src/repository"
	"online-food/src/usecase"

	// "log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	r := server.InitServer()

	// godotenv.Load()
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println("Error loading .env file:", err)
	// 	return
	// }

	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbParams := os.Getenv("DB_PARAMS")

	// Construct the connection string
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?%s",
		dbUsername, dbPassword, dbHost, dbPort, dbName, dbParams,
	)

	fmt.Println("connectionString>>>>>>", connectionString)

	postgreConfig := properties.PostgreConfig{
		DatabaseURL: connectionString,
	}

	db := db.InitPostgreDB(postgreConfig)

	//run migrations
	m, err := migrate.New(os.Getenv("MIGRATION_PATH"), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error creating migration instance: ", err)
	}

	//Run the migration up to the latest version
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Error applying migrations:", err)
	}

	fmt.Println("Migration successfully applied")

	helper := helpers.NewHelper()

	// MIDDLEWARE
	middleware := middleware.NewMiddleware(helper)

	// REPOSITORY
	userRepository := repository.NewUserRepository(db)
	merchantRepository := repository.NewMerchantRepository(db)

	// USECASE
	authUsecase := usecase.NewAuthUsecase(userRepository, helper)
	merchantUsecase := usecase.NewMerchantUsecase(merchantRepository, helper)

	// HANDLER
	authHandler := handler.NewAuthHandler(authUsecase)
	imageHandler := handler.NewImageHandler()
	merchantHandler := handler.NewMerchantHandler(merchantUsecase)

	// ROUTE
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.POST("/admin/register", authHandler.Register)
	r.POST("/admin/login", authHandler.Login)
	r.POST("/user/register", authHandler.Register)
	r.POST("/user/login", authHandler.Login)

	authorized := r.Group("")
	authorized.Use(middleware.AuthMiddleware)

	// IT user only
	adminAuthorized := authorized.Group("")
	adminAuthorized.Use(middleware.RoleMiddleware("admin"))
	r.POST("/admin/merchants", merchantHandler.CreateMerchant)
	r.GET("/admin/merchants", merchantHandler.GetMerchants)

	// upload image
	authorized.POST("/image", imageHandler.UploadImage)

	r.Run()
}
