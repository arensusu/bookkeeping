package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"firstapp/controller"
	"firstapp/database"
	"firstapp/middleware"
	"firstapp/model"
)

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{}, &model.Category{})
	database.Database.AutoMigrate(&model.Detail{})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	loadDatabase()

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(config))

	authRoutes := router.Group("auth")
	authRoutes.POST("/register", controller.Register)
	authRoutes.POST("/login", controller.Login)
	authRoutes.GET("/user", middleware.Verify(), controller.GetUser)
	authRoutes.DELETE("/user", middleware.Verify(), controller.DeleteUser)

	apiRoutes := router.Group("api")
	apiRoutes.GET("/categories", controller.GetCategories)
	apiRoutes.POST("/categories", controller.CreateCategory)
	apiRoutes.GET("/categories/:id", controller.GetCategory)

	details := apiRoutes.Group("/details")
	details.Use(middleware.Verify())
	details.GET("", controller.GetDetails)
	details.POST("", controller.CreateDetail)
	details.GET("/:id", controller.GetDetail)
	details.DELETE("/:id", controller.DeleteDetail)
	router.Run(":8080")
}
