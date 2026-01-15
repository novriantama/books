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
	bookRepo := repository.NewBookRepository(database)
	authorRepo := repository.NewAuthorRepository(database)
	pubRepo := repository.NewPublisherRepository(database)

	userService := services.NewUserService(userRepo)
	bookService := services.NewBookService(bookRepo)
	authorService := services.NewAuthorService(authorRepo)
	pubService := services.NewPublisherService(pubRepo)

	userHandler := handlers.NewUserHandler(userService)
	bookHandler := handlers.NewBookHandler(bookService)
	authorHandler := handlers.NewAuthorHandler(authorService)
	pubHandler := handlers.NewPublisherHandler(pubService)

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

		library := api.Group("/library")
		library.Use(middleware.AuthMiddleware())
		{
			library.POST("/books", bookHandler.Create)
			library.GET("/books", bookHandler.List)
			library.GET("/books/:id", bookHandler.Get)

			library.POST("/authors", authorHandler.Create)
			library.GET("/authors", authorHandler.List)
			library.PUT("/authors/:id", authorHandler.Update)
			library.DELETE("/authors/:id", authorHandler.Delete)

			library.POST("/publishers", pubHandler.Create)
			library.GET("/publishers", pubHandler.List)
			library.PUT("/publishers/:id", pubHandler.Update)
			library.DELETE("/publishers/:id", pubHandler.Delete)
		}
	}

	r.Run(":8080")
}
