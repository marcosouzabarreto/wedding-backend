package routes

import (
	"github.com/gin-gonic/gin"
	"wedding-backend/handlers"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/login", handlers.Login)
}
