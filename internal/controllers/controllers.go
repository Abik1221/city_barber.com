package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"city_barber.com/internal/models"
	"city_barber.com/internal/database"
)

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save user to the database
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user})
}

// LoginUser handles user login
func LoginUser(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authenticate user (pseudo-code)
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check password (pseudo-code)
	if err := helpers.ComparePassword(user.PasswordHash, input.Password); err != nil {
	    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
	    return
	}

	// Generate token (pseudo-code)
	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
	    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
	    return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": "dummy-token"})
}

// GetUserProfile returns the authenticated user's profile
func GetUserProfile(c *gin.Context) {
	userID := c.GetUint("userID") // Assuming userID is set in the middleware

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateUserProfile updates the authenticated user's profile
func UpdateUserProfile(c *gin.Context) {
	userID := c.GetUint("userID")

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update user fields
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.PhoneNumber = input.PhoneNumber
	user.Gender = input.Gender

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

// DeleteUser deletes the authenticated user's account
func DeleteUser(c *gin.Context) {
	userID := c.GetUint("userID")

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// GetAllBarbers returns a list of all barbers
func GetAllBarbers(c *gin.Context) {
	var barbers []models.Barber
	if err := database.DB.Find(&barbers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch barbers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"barbers": barbers})
}

// GetBarberByID returns a specific barber by ID
func GetBarberByID(c *gin.Context) {
	barberID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid barber ID"})
		return
	}

	var barber models.Barber
	if err := database.DB.First(&barber, barberID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Barber not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"barber": barber})
}

// CreateBarber creates a new barber (admin-only)
func CreateBarber(c *gin.Context) {
	var barber models.Barber
	if err := c.ShouldBindJSON(&barber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&barber).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create barber"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Barber created successfully", "barber": barber})
}

// UpdateBarber updates a barber's details (admin-only)
func UpdateBarber(c *gin.Context) {
	barberID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid barber ID"})
		return
	}

	var input models.Barber
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var barber models.Barber
	if err := database.DB.First(&barber, barberID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Barber not found"})
		return
	}

	// Update barber fields
	barber.FirstName = input.FirstName
	barber.LastName = input.LastName
	barber.ProfilePicture = input.ProfilePicture
	barber.Bio = input.Bio
	barber.SocialMediaLink = input.SocialMediaLink

	if err := database.DB.Save(&barber).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update barber"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Barber updated successfully", "barber": barber})
}

// DeleteBarber deletes a barber (admin-only)
func DeleteBarber(c *gin.Context) {
	barberID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid barber ID"})
		return
	}

	var barber models.Barber
	if err := database.DB.First(&barber, barberID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Barber not found"})
		return
	}

	if err := database.DB.Delete(&barber).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete barber"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Barber deleted successfully"})
}