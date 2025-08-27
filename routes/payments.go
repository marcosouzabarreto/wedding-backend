package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wedding-backend/handlers"
)

func PaymentRoutes(router *gin.Engine, db *gorm.DB) {
	payment := router.Group("/payments")
	{
		payment.POST("/create-preference", handlers.CreatePayment(db))
	}
}
