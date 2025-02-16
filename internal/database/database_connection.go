package database

import (
	"fmt"
	"log"

	"city_barber.com/configs"
	"city_barber.com/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection and performs auto-migration
func InitDB(dbUser, dbPassword, dbHost, dbPort, dbName string) error {
	// Create the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open a connection to the database
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Connected to the database")

	// Perform auto-migration
	if err := DB.AutoMigrate(
		&models.User{},
		&models.Barber{},
		&models.Shop{},
		&models.Service{},
		&models.Booking{},
		&models.Payment{},
		&models.PromoCode{},
		&models.Admin{},
	); err != nil {
		return fmt.Errorf("failed to auto-migrate database: %v", err)
	}

	log.Println("Database migration completed")

	return nil
}