package main

import (
	"fmt"
	"health-record/app/server"
	"health-record/db"
	"health-record/helpers"
	"health-record/model/properties"
	"health-record/src/handler"
	"health-record/src/middleware"
	"health-record/src/repository"
	"health-record/src/usecase"

	// "log"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/golang-migrate/migrate/v4"
	// _ "github.com/golang-migrate/migrate/v4/database/postgres"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
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

	// //run migrations
	// m, err := migrate.New(os.Getenv("MIGRATION_PATH"), os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatal("Error creating migration instance: ", err)
	// }

	// //Run the migration up to the latest version
	// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	// 	log.Fatal("Error applying migrations:", err)
	// }

	fmt.Println("Migration successfully applied")

	helper := helpers.NewHelper()

	// MIDDLEWARE
	middleware := middleware.NewMiddleware(helper)

	// REPOSITORY
	userRepository := repository.NewUserRepository(db)
	nurseRepository := repository.NewNurseRepository(db)

	// USECASE
	authUsecase := usecase.NewAuthUsecase(userRepository, helper)
	nurseUscase := usecase.NewNurseUsecase(nurseRepository)

	// HANDLER
	authHandler := handler.NewAuthHandler(authUsecase)
	nurseHandler := handler.NewNurseHandler(nurseUscase)
	imageHandler := handler.NewImageHandler()

	// ROUTE
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.POST("/v1/user/it/register", authHandler.Register)
	r.POST("/v1/user/it/login", authHandler.Login)
	r.POST("/v1/user/nurse/login", authHandler.LoginNurse)

	authorized := r.Group("")
	authorized.Use(middleware.AuthMiddleware)

	// IT user only
	itAuthorized := authorized.Group("")
	itAuthorized.Use(middleware.RoleMiddleware("it"))
	itAuthorized.POST("/v1/user/nurse/register", nurseHandler.RegisterNurse)
	itAuthorized.GET("/v1/user", nurseHandler.GetUsers)
	itAuthorized.PUT("/v1/user/nurse/:userId", nurseHandler.UpdateNurse)
	itAuthorized.DELETE("/v1/user/nurse/:userId", nurseHandler.DeleteNurse)
	itAuthorized.POST("/v1/user/nurse/:userId/access", nurseHandler.AddAccess)

	// upload image
	authorized.POST("/v1/image", imageHandler.UploadImage)

	r.Run()
}
