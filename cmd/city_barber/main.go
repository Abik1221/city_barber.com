package main

import (
	"log"
	"github.com/abik1221/city_barber.com/configs"
	"github.com/abik1221/city_barber.com/internal/database"
	"github.com/abik1221/city_barber.com/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config := configs.LoadConfig()

	if err := database.InitDB(config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}


	router := gin.Default()


	routes.SetupRoutes(router)


	port := config.Port
	if port == "" {
		port = "8080" 
	}

	log.Printf("Server is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}