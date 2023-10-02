package main

import (
	"bookkeeping-backend/database"
	"bookkeeping-backend/model"
	"fmt"
	"time"
)

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{}, &model.Category{})
	database.Database.AutoMigrate(&model.Detail{})
}

func init() {
	fmt.Println(time.Now())
}

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic("Error loading .env file")
	// }

	// loadDatabase()

	// router := gin.Default()
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"https://localhost:3000", "https://arensusu.github.io"}
	// config.AllowHeaders = []string{"Authorization", "Content-Type"}
	// router.Use(cors.New(config))

	// authRoutes := router.Group("/auth")
	// authRoutes.POST("/register", controller.Register)
	// authRoutes.POST("/login", controller.Login)
	// authRoutes.GET("/user", middleware.Verify(), controller.GetUser)
	// authRoutes.DELETE("/user", middleware.Verify(), controller.DeleteUser)

	// apiRoutes := router.Group("/api")

	// categoryRoutes := apiRoutes.Group("/categories")
	// categoryRoutes.GET("", controller.GetAllCategories)
	// categoryRoutes.POST("", controller.CreateCategory)
	// categoryRoutes.GET("/:id", controller.GetCategory)
	// categoryRoutes.DELETE("/:id", controller.DeleteCategory)
	// categoryRoutes.PUT("/:id", controller.UpdateCategory)

	// detailRoutes := apiRoutes.Group("/details")
	// detailRoutes.Use(middleware.Verify())
	// detailRoutes.GET("", controller.GetDetails)
	// detailRoutes.POST("", controller.CreateDetail)
	// detailRoutes.GET("/:id", controller.GetDetail)
	// detailRoutes.DELETE("/:id", controller.DeleteDetail)
	// router.Run(":8080")
	//router.RunTLS(":8080", "./certs/server.crt", "./certs/server.key")
	fmt.Println(time.Now())
}
