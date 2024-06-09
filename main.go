package main

import (
	"fmt"
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
	// "github.com/golang-migrate/migrate/v4"
	// _ "github.com/golang-migrate/migrate/v4/database/postgres"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	r := server.InitServer()

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

	// //run migrations
	// m, err := migrate.New(os.Getenv("MIGRATION_PATH"), os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatal("Error creating migration instance: ", err)
	// }

	// //Run the migration up to the latest version
	// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	// 	log.Fatal("Error applying migrations:", err)
	// }

	// fmt.Println("Migration successfully applied")

	authHelper := helpers.NewAuthHelper()

	// MIDDLEWARE
	middleware := middleware.NewMiddleware(authHelper)

	// REPOSITORY
	userRepository := repository.NewUserRepository(db)
	merchantRepository := repository.NewMerchantRepository(db)
	orderRepository := repository.NewOrderRepository(db)

	// USECASE
	merchantUsecase := usecase.NewMerchantUsecase(merchantRepository)
	authUsecase := usecase.NewAuthUsecase(userRepository, authHelper)
	purchaseUseCase := usecase.NewPurchaseUsecase(merchantRepository, orderRepository)

	// HANDLER
	authHandler := handler.NewAuthHandler(authUsecase)
	imageHandler := handler.NewImageHandler()
	merchantHandler := handler.NewMerchantHandler(merchantUsecase)
	purchaseHandler := handler.NewPurchaseHandler(purchaseUseCase, merchantUsecase)

	// ROUTE
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.POST("/admin/register", authHandler.Register)
	r.POST("/admin/login", authHandler.Login)
	r.POST("/users/register", authHandler.Register)
	r.POST("/users/login", authHandler.Login)

	authorized := r.Group("")
	authorized.Use(middleware.AuthMiddleware)

	// admin only
	adminAuthorized := authorized.Group("")
	adminAuthorized.Use(middleware.RoleMiddleware("admin"))
	adminAuthorized.POST("/admin/merchants", merchantHandler.CreateMerchant)
	adminAuthorized.GET("/admin/merchants", merchantHandler.GetMerchants)

	adminAuthorized.POST("/admin/merchants/:id/items", merchantHandler.CreateMerchantItem)
	adminAuthorized.GET("/admin/merchants/:id/items", merchantHandler.GetMerchantItems)

	// upload image
	authorized.POST("/image", imageHandler.UploadImage)

	// purchase
	// authorized.GET("/merchants/nearby/:lat/:long", purchaseHandler.GetNearbyMerchants)
	authorized.GET("/merchants/nearby/:latLong", purchaseHandler.GetNearbyMerchants)
	authorized.POST("/users/estimate", purchaseHandler.CreateEstimation)
	authorized.POST("/users/orders", purchaseHandler.CreateOrder)
	authorized.GET("/users/orders", purchaseHandler.GetOrders)

	r.Run()
}
  