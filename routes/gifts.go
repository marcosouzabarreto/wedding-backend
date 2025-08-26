package routes

import (
	"wedding-backend/handlers"
	"wedding-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupGiftRoutes(r *gin.Engine, db *gorm.DB) {
	gift := r.Group("/gifts")
	giftService := services.NewGiftService(db)
	giftRoutes := handlers.NewGiftHandlers(giftService)
	{
		gift.GET("/", giftRoutes.GetAll)
		gift.POST("/", giftRoutes.Create)
		gift.GET("/:id", giftRoutes.GetByID)
		gift.PUT("/:id", giftRoutes.Update)
		gift.DELETE("/:id", giftRoutes.Delete)
	}
}
