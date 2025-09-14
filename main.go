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
		userAuthGroup.GET("/orders", controllers.GetMyOrdersHandler())
		userAuthGroup.POST("/orders/:orderId/cancel", controllers.CancelOrderHandler())
		userAuthGroup.POST("/orders/:orderId/review", controllers.LeaveReviewHandler())

	}

	router.Run(":8080")
}
