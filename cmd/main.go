package main

import (
	"fmt"
	"log"
	"user-services/internal/delivery/http/handler"
	"user-services/internal/domain"
	"user-services/internal/repository"
	"user-services/internal/usecase"
	"user-services/pkg/config"
	"user-services/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Database
	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Auto Migrate
	db.AutoMigrate(&domain.User{})

	// Initialize Repository
	userRepo := repository.NewUserRepository(db)

	// Initialize Usecase
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Initialize Handler
	userHandler := handler.NewUserHandler(userUsecase)

	// Initialize Router
	r := gin.Default()

	// Routes
	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/", userHandler.Create)
			users.GET("/:id", userHandler.GetByID)
			users.GET("/email/:email", userHandler.GetByEmail)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
		}
	}

	// Start Server
	if err := r.Run(fmt.Sprintf(":%s", cfg.AppPort)); err != nil {
		log.Fatal(err)
	}
}
