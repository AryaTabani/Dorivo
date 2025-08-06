package handler

import (
	"net/http"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/controllers"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	db.InitDB()

	router = gin.Default()
	router.GET("/tenant/:tenantId", controllers.GetTenantConfigHandler())
	router.POST("/register", controllers.RegisterHandler())
	router.POST("/login", controllers.LoginHandler())

}
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
