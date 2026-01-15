package main

import (
	"books/db"
	"books/pkg/config"
	"books/pkg/handlers"
	"books/pkg/middleware"
	"books/pkg/repository"
	"books/pkg/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load Env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	// Connect DB
	database, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Run Migration & Seed
	db.Migrate(database)
	db.Seed(database)

	// Dependency Injection
	userRepo := repository.NewUserRepository(database)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Router
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)

		// Protected Routes
		protected := api.Group("/user")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/profile", userHandler.Profile)
		}
	}

	r.Run(":8080")
}
