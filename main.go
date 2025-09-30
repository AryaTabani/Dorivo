package main

import (
	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/controllers"
	"github.com/AryaTabani/Dorivo/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	router := gin.Default()

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
	}

	router.Run(":8080")
}
