package routes

import (
	"wedding-backend/handlers"
	"wedding-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserGiftRoutes(r *gin.Engine, db *gorm.DB) {
	userGift := r.Group("/user_gifts")
	userGiftService := services.NewUserGiftService(db)
	userGiftRoutes := handlers.NewUserGiftHandlers(userGiftService)
	{
		userGift.GET("/", userGiftRoutes.GetAll)
		userGift.POST("/", userGiftRoutes.Create)
	}
}
