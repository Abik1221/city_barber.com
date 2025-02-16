package routes

import (
	"github.com/gin-gonic/gin"
	"city_barber.com/internal/controllers"
	"city_barber.com/internal/middleware"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes (no authentication required)
	public := router.Group("/api")
	{
		// User routes
		public.POST("/register", controllers.RegisterUser)
		public.POST("/login", controllers.LoginUser)

		// Barber routes
		public.GET("/barbers", controllers.GetAllBarbers)
		public.GET("/barbers/:id", controllers.GetBarberByID)

		// Shop routes
		public.GET("/shops", controllers.GetAllShops)
		public.GET("/shops/:id", controllers.GetShopByID)

		// Service routes
		public.GET("/services", controllers.GetAllServices)
		public.GET("/services/:id", controllers.GetServiceByID)
	}

	// Protected routes (authentication required)
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// User routes
		protected.GET("/user", controllers.GetUserProfile)
		protected.PUT("/user", controllers.UpdateUserProfile)
		protected.DELETE("/user", controllers.DeleteUser)

		// Barber routes
		protected.POST("/barbers", controllers.CreateBarber)
		protected.PUT("/barbers/:id", controllers.UpdateBarber)
		protected.DELETE("/barbers/:id", controllers.DeleteBarber)

		// Booking routes
		protected.POST("/bookings", controllers.CreateBooking)
		protected.GET("/bookings", controllers.GetUserBookings)
		protected.GET("/bookings/:id", controllers.GetBookingByID)
		protected.PUT("/bookings/:id", controllers.UpdateBooking)
		protected.DELETE("/bookings/:id", controllers.CancelBooking)

		// Payment routes
		protected.POST("/payments", controllers.CreatePayment)
		protected.GET("/payments/:id", controllers.GetPaymentByID)

		// Admin routes
		protected.POST("/promocodes", controllers.CreatePromoCode)
		protected.GET("/promocodes", controllers.GetAllPromoCodes)
		protected.DELETE("/promocodes/:id", controllers.DeletePromoCode)
	}
}