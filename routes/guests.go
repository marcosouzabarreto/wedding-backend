package routes

import (
	"wedding-backend/handlers"
	"wedding-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupGuestRoutes(r *gin.Engine, db *gorm.DB) {
	guest := r.Group("/guests")
	guestService := services.NewGuestService(db)
	guestRoutes := handlers.NewGuestHandlers(guestService)
	{
		guest.GET("/", guestRoutes.GetAll)
		guest.POST("/", guestRoutes.Create)
		guest.GET("/:id", guestRoutes.GetByID)
	}
}
