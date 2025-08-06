package main

import (
	db "example.com/m/v2/DB"
	"example.com/m/v2/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	router := gin.Default()
	router.GET("/config/:tenantId", controllers.GetTenantConfigHandler())
	router.POST("/register", controllers.RegisterHandler())
	router.POST("/login", controllers.LoginHandler())

	router.Run(":8080")
}
