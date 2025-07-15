package main

import (
	"task-manager/controllers"
	"task-manager/database"
	"task-manager/middleware"

	_ "task-manager/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Task Manager API
// @version 1.0
// @description This is a simple task manager API
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
// @description JWT Authorization header using the Bearer scheme. Enter your token only (without Bearer prefix).
// @Security Bearer
func main() {
	database.Connect()
	r := gin.Default()

	// Configure Swagger
	config := &ginSwagger.Config{
		URL:                  "http://localhost:8080/swagger/doc.json", // The url pointing to API definition
		DeepLinking:          true,
		DocExpansion:         "list",
		PersistAuthorization: true,
	}

	// Public routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// Protected routes
	api := r.Group("/")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/tasks", controllers.GetTasks)
		api.GET("/tasks/:id", controllers.GetTask)
		api.POST("/tasks", controllers.CreateTask)
		api.PUT("/tasks/:id", controllers.UpdateTask)
		api.DELETE("/tasks/:id", controllers.DeleteTask)
	}

	// Swagger documentation - no authentication required to view the docs
	r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))

	r.Run() // Listen and serve on 0.0.0.0:8080
}
