package main

import (
	"log"
	"city_barber.com/configs"
	"city_barber.com/internal/database"
	"city_barber.com/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config := configs.LoadConfig()

	// Initialize the database
	if err := database.InitDB(config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Set up Gin
	router := gin.Default()

	// Set up routes
	routes.SetupRoutes(router)

	// Start the server
	port := config.Port
	if port == "" {
		port = "8080" // Default port
	}

	log.Printf("Server is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}