package routes

import (
	"github.com/abik1221/city_barber.com/internal/database"
	"github.com/abik1221/city_barber.com/internal/controllers"
	"github.com/abik1221/city_barber.com/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(incomig_routes *gin.Engine) {
	authController := controllers.NewAuthController(services.NewAuthService(database.DB))
	googleAuthController := controllers.NewGoogleAuthController(services.NewGoogleAuthService(database.DB))


	public := incomig_routes.Group("/api")
	{
		public.POST("/login", authController.Login)
		public.POST("/forgot-password", authController.ForgotPassword)
		public.GET("/auth/google", googleAuthController.GoogleLogin)
	}

	protected := incomig_routes.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/user", controllers.GetUserProfile)
	}
}