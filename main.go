package main

import (
	"log"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/controllers"
	_ "github.com/AryaTabani/Dorivo/docs"
	"github.com/AryaTabani/Dorivo/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Dorivo Multi-Tenant API
// @version 1.0
// @description This is the backend API for the Dorivo multi-tenant ordering system.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and a JWT token.

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}
	db.InitDB()
	db.InitRedis()
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/tenant/:tenantId", controllers.GetTenantConfigHandler())
	router.POST("/:tenantId/register", controllers.RegisterHandler())
	router.POST("/:tenantId/login", controllers.LoginHandler())
	router.GET("/:tenantId/faqs", controllers.GetFAQsHandler())
	router.GET("/:tenantId/products", controllers.SearchProductsHandler())
	router.GET("/:tenantId/tags", controllers.GetTagsHandler())
	router.GET("/:tenantId/products/:productId", controllers.GetProductDetailsHandler())
	router.GET("/:tenantId/products/bestsellers", controllers.GetBestSellersHandler())
	router.GET("/:tenantId/products/featured", controllers.GetFeaturedProductHandler())
	router.GET("/:tenantId/products/recommended", controllers.GetRecommendedProductsHandler())
	router.POST("/superadmin/login", controllers.SuperAdminLoginHandler())

	userAuthGroup := router.Group("/")
	userAuthGroup.Use(middleware.AuthMiddleware())
	{
		userAuthGroup.GET("/profile", controllers.GetProfileHandler())
		userAuthGroup.PUT("/profile", controllers.UpdateProfileHandler())

		userAuthGroup.GET("/addresses", controllers.GetAddressesHandler())
		userAuthGroup.POST("/addresses", controllers.AddAddressHandler())
		userAuthGroup.DELETE("/addresses/:addressId", controllers.DeleteAddressHandler())

		userAuthGroup.GET("/payment-methods", controllers.GetPaymentMethodsHandler())
		userAuthGroup.POST("/payment-methods", controllers.AddPaymentMethodHandler())
		userAuthGroup.DELETE("/payment-methods/:methodId", controllers.DeletePaymentMethodHandler())

		userAuthGroup.GET("/orders", controllers.GetMyOrdersHandler())
		userAuthGroup.POST("/orders/:orderId/cancel", controllers.CancelOrderHandler())
		userAuthGroup.POST("/orders/:orderId/review", controllers.LeaveReviewHandler())

		userAuthGroup.GET("/profile/notification-settings", controllers.GetNotificationsSettingHandler())
		userAuthGroup.PUT("/profile/notification-settings", controllers.UpdateNotificationSettingsHandler())
		userAuthGroup.PUT("/profile/change-password", controllers.ChangePasswordHandler())
		userAuthGroup.DELETE("/profile", controllers.DeleteAccountHandler())

		userAuthGroup.GET("/notifications", controllers.GetNotificationsHandler())
		userAuthGroup.PUT("/notifications/read", controllers.MarkReadHandler())

		userAuthGroup.GET("/cart", controllers.GetCartHandler())
		userAuthGroup.POST("/cart/items", controllers.AddToCartHandler())
		userAuthGroup.PUT("/cart/items/:itemId", controllers.UpdateCartItemHandler())
		userAuthGroup.DELETE("/cart/items/:itemId", controllers.RemoveCartItemHandler())

		userAuthGroup.GET("/favorites", controllers.GetFavoritesHandler())
		userAuthGroup.POST("/products/:productId/favorite", controllers.AddToFavoritesHandler())
		userAuthGroup.DELETE("/products/:productId/favorite", controllers.RemoveFromFavoritesHandler())
	}

	adminGroup := router.Group("/:tenantId/admin")
	adminGroup.Use(middleware.AdminAuthMiddleware())
	{
		adminGroup.POST("/products", controllers.CreateProductHandler())
		adminGroup.PUT("/products/:productId", controllers.UpdateProductHandler())
		adminGroup.DELETE("/products/:productId", controllers.DeleteProductHandler())
		adminGroup.PUT("/config", controllers.UpdateTenantConfigHandler())

		adminGroup.GET("/orders", controllers.GetTenantOrdersHandler())
		adminGroup.PUT("/orders/:orderId/status", controllers.UpdateOrderStatusHandler())

		adminGroup.GET("/customers", controllers.GetTenantCustomersHandler())
		adminGroup.GET("/dashboard/stats", controllers.GetDashboardStatsHandler())

	}
	superAdminGroup := router.Group("/superadmin")
	superAdminGroup.Use(middleware.SuperAdminAuthMiddleware())
	{
		superAdminGroup.GET("/tenants", controllers.GetAllTenantsHandler())
		superAdminGroup.POST("/tenants", controllers.CreateTenantHandler())
		superAdminGroup.DELETE("/tenants/:tenantId", controllers.DeleteTenantHandler())
	}
	router.Run(":8080")
}
