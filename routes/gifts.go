package routes

import (
	"wedding-backend/handlers"
	"wedding-backend/middleware"
	"wedding-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupGiftRoutes(r *gin.Engine, db *gorm.DB) {
	giftService := services.NewGiftService(db)
	giftRoutes := handlers.NewGiftHandlers(giftService)

	giftPublic := r.Group("/gifts")
	{
		giftPublic.GET("/", giftRoutes.GetAll)
		giftPublic.GET("/:id", giftRoutes.GetByID)
	}

	giftProtected := r.Group("/gifts")
	giftProtected.Use(middleware.AuthMiddleware())
	{
		giftProtected.POST("/", giftRoutes.Create)
		giftProtected.PUT("/:id", giftRoutes.Update)
		giftProtected.DELETE("/:id", giftRoutes.Delete)
	}
}
