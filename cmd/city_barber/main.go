package main

import (
	"log"
	"os"

	"city_barber.com/configs"
	"city_barber.com/internal/database"
	"city_barber.com/internal/routes"
	"github.com/gin-gonic/gin"

	"github.com/joho/go.env"
)

func main() {
	// Load environment variables or configurations
	configs.LoadConfig()

	// Initialize the database connection
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Set up Gin
	router := gin.Default()

	// Set up routes
	routes.SetupRoutes(router)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	log.Printf("Server is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
