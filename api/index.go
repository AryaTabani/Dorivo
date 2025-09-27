package api

import (
	"net/http"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/controllers"
	"github.com/AryaTabani/Dorivo/middleware"
	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	db.InitDB()

	router := gin.Default()

	router.GET("/tenant/:tenantId", controllers.GetTenantConfigHandler())
	router.POST("/:tenantId/register", controllers.RegisterHandler())
	router.POST("/:tenantId/login", controllers.LoginHandler())
	router.GET("/:tenantId/faqs", controllers.GetFAQsHandler())

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
	}
	router.ServeHTTP(w, r)
}
