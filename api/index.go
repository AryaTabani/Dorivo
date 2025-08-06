package handler

import (
	"net/http"

	db "example.com/m/v2/DB"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	db.InitDB()

	router = gin.Default()

}
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
