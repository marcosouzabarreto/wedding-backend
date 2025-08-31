package routes

import (
	"wedding-backend/handlers"
	"wedding-backend/middleware"
	"wedding-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupGuestRoutes(r *gin.Engine, db *gorm.DB) {
	guestService := services.NewGuestService(db)
	guestRoutes := handlers.NewGuestHandlers(guestService)

	guestPublic := r.Group("/guests")
	{
		guestPublic.GET("/:id", guestRoutes.GetByID)
	}

	guestProtected := r.Group("/guests")
	guestProtected.Use(middleware.AuthMiddleware())
	{
		guestProtected.GET("/", guestRoutes.GetAll)
		guestProtected.POST("/", guestRoutes.Create)
		guestProtected.PUT("/:id", guestRoutes.Update)
		guestProtected.DELETE("/:id", guestRoutes.Delete)
	}
}
