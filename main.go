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
	}

	router.Run(":8080")
}
