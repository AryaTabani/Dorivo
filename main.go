package main

import (
	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	router := gin.Default()
	router.GET("/tenant/:tenantId", controllers.GetTenantConfigHandler())
	router.POST("/register", controllers.RegisterHandler())
	router.POST("/login", controllers.LoginHandler())

	router.Run(":8080")
}
